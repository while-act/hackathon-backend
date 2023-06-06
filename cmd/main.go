package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	_ "github.com/while-act/hackathon-backend/docs"
	"github.com/while-act/hackathon-backend/ent"
	"github.com/while-act/hackathon-backend/internal/controller"
	"github.com/while-act/hackathon-backend/internal/repo/postgres"
	redisRepo "github.com/while-act/hackathon-backend/internal/repo/redis"
	"github.com/while-act/hackathon-backend/internal/service"
	"github.com/while-act/hackathon-backend/pkg/client/email"
	"github.com/while-act/hackathon-backend/pkg/client/postgresql"
	redisInit "github.com/while-act/hackathon-backend/pkg/client/redis"
	"github.com/while-act/hackathon-backend/pkg/conf"
	"github.com/while-act/hackathon-backend/pkg/log"
	"github.com/while-act/hackathon-backend/pkg/middleware/bind"
	"github.com/while-act/hackathon-backend/pkg/middleware/errs"
	"github.com/while-act/hackathon-backend/pkg/middleware/query"
	"github.com/while-act/hackathon-backend/pkg/middleware/session"
	"net/http"
	"net/smtp"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// @title While.act API
// @version 1.0
// @description It's an API interacting with While.act using Golang
// @accept application/json
// @produce application/json
// @schemes http
// @BasePath /api

// @contact.name Contact us
// @contact.url https://github.com/while-act/hackathon-backend/issues/new/choose
// @contact.email  matvey-sizov@mail.ru

// @securityDefinitions.apiKey  ApiKeyAuth
// @in header
// @name session_id
// @host 37.230.195.26:3000

// @sessions.docs.description Authorization, registration and authentication
func main() {
	cfg := conf.GetConfig()

	pClient, rClient, mailClient := getClients(cfg)

	h := initHandler(pClient, rClient, mailClient, cfg)
	r := gin.New()

	h.InitRoutes(createSetter(r, pClient, rClient, cfg, mailClient != nil))

	run(cfg.Listen.Port, r, pClient, rClient, mailClient)
}

// run the Server with graceful shutdown
func run(port int, r *gin.Engine, pClient *ent.Client, rClient *redis.Client, mailClient *smtp.Client) {
	srv := &http.Server{
		Addr:           fmt.Sprintf(":%d", port),
		Handler:        r,
		MaxHeaderBytes: 1 << 20, // 1 MB
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.WithErr(err).Fatalf("error occurred while running http server")
		}
	}()
	log.Infof("Server Started On Port %d", port)

	<-quit

	log.Info("Server Shutting Down ...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.WithErr(err).Fatal("Server Shutdown Failed")
	}

	if err := rClient.Close(); err != nil {
		log.WithErr(err).Fatal("Redis Connection Shutdown Failed")
	}

	if err := pClient.Close(); err != nil {
		log.WithErr(err).Fatal("PostgreSQL Connection Shutdown Failed")
	}

	log.Info("Server Exited Properly")
	if err := mailClient.Quit(); err != nil {
		log.WithErr(err).Fatal("Email Connection Shutdown Failed")
	}
}

func getClients(cfg *conf.Config) (*ent.Client, *redis.Client, *smtp.Client) {
	pClient := postgresql.Open(cfg.DB.Postgres.Username, cfg.DB.Postgres.Password,
		cfg.DB.Postgres.Host, cfg.DB.Postgres.Port, cfg.DB.Postgres.DBName)

	rClient := redisInit.Open(cfg.DB.Redis.Host, cfg.DB.Redis.Port, cfg.DB.Redis.DbId)

	mailClient := email.Open(cfg.Email.User, cfg.Email.Password, cfg.Email.Host, cfg.Email.Port)

	return pClient, rClient, mailClient
}

func initHandler(pClient *ent.Client, rClient *redis.Client, mailClient *smtp.Client, cfg *conf.Config) *controller.Handler {
	pUser := postgres.NewUserStorage(pClient.User)
	pComp := postgres.NewCompanyStorage(pClient.Company)
	pHist := postgres.NewHistoryStorage(pClient.History)
	pInd := postgres.NewIndustryStorage(pClient.Industry)
	pDist := postgres.NewDistrictStorage(pClient.District)
	pTax := postgres.NewTaxStorage(pClient.TaxationSystem)
	pBus := postgres.NewBusinessStorage(pClient.BusinessActivity)
	rConn := redisRepo.NewRClient(rClient)

	auth := service.NewAuthService(pUser, rConn)
	history := service.NewHistoryService(pHist)
	industry := service.NewIndustryService(pInd)
	district := service.NewDistrictService(pDist)
	taxationSystem := service.NewTaxService(pTax)
	business := service.NewBusinessService(pBus)
	user := service.NewUserService(pUser, rConn)
	pdf := service.NewPDF(cfg.TemplatePath)
	mail := service.NewEmailSender(mailClient)
	company := service.NewCompanyService(pComp)

	return controller.NewHandler(
		user,
		company,
		auth,
		history,
		industry,
		district,
		taxationSystem,
		business,
		pdf,
		session.NewAuth(auth, cfg),
		mail,
	)
}

func createSetter(r *gin.Engine, pClient *ent.Client, rClient *redis.Client, cfg *conf.Config, mailSet bool) *controller.Setter {
	pUser := postgres.NewUserStorage(pClient.User)
	rConn := redisRepo.NewRClient(rClient)

	auth := service.NewAuthService(pUser, rConn)

	return controller.NewSetter(
		r,
		bind.NewValidator(),
		errs.NewErrHandler(),
		query.NewQueryHandler(),
		session.NewAuth(auth, cfg),
		cfg.Listen.MainPath,
		mailSet,
	)
}
