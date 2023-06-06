package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/while-act/hackathon-backend/internal/controller/dao"
	"github.com/while-act/hackathon-backend/pkg/middleware/bind"
	"github.com/while-act/hackathon-backend/pkg/middleware/session"
)

type ErrHandler interface {
	HandleError(handler func(*gin.Context) error) gin.HandlerFunc
}

type SessionHandler interface {
	Session(handler func(ctx *gin.Context, session *dao.Session) error) func(c *gin.Context) error
	SessionFunc(c *gin.Context) (*dao.Session, error)
}

type QueryHandler interface {
	HandleQueries() gin.HandlerFunc
}

type Setter struct {
	r        *gin.Engine
	valid    *validator.Validate
	erh      ErrHandler
	qh       QueryHandler
	sess     SessionHandler
	mainPath string
	mailSet  bool
}

func NewSetter(r *gin.Engine, valid *validator.Validate, erh ErrHandler, qh QueryHandler, sess SessionHandler, mainPath string, mailSet bool) *Setter {
	return &Setter{r: r, valid: valid, erh: erh, qh: qh, sess: sess, mainPath: mainPath, mailSet: mailSet}
}

func (h *Handler) InitRoutes(s *Setter) {
	s.r.Use(s.cors, s.qh.HandleQueries(), gin.Recovery())
	rg := s.r.Group(s.mainPath)

	calc := rg.Group("/calc")
	{
		calc.GET("/:industry", s.erh.HandleError(bind.HandleParam(h.getIndustryInfo, "industry", "required", s.valid)))
		calc.POST("/save", s.erh.HandleError(session.HandleBody(h.saveCalcData, s.sess.SessionFunc, s.valid)))
		calc.POST("", s.erh.HandleError(bind.HandleBody(h.calcData, s.valid)))
	}

	docs := rg.Group("/docs")
	{
		docs.GET("/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	auth := rg.Group("/auth")
	{
		auth.POST("/sign-up", s.erh.HandleError(bind.HandleBody(h.signUp, s.valid)))
		auth.POST("/sign-in", s.erh.HandleError(bind.HandleBody(h.signIn, s.valid)))

		sess := auth.Group("/session")
		{
			sess.GET("", s.erh.HandleError(s.sess.Session(h.getMe)))
			sess.DELETE("", s.erh.HandleError(s.sess.Session(h.signOut)))
		}
	}

	user := rg.Group("/user")
	{
		user.PATCH("", s.erh.HandleError(session.HandleBody(h.updateMe, s.sess.SessionFunc, s.valid)))
		user.PATCH("/email", s.erh.HandleError(session.HandleBody(h.updateEmail, s.sess.SessionFunc, s.valid)))
		user.PATCH("/password", s.erh.HandleError(bind.HandleBody(h.updatePassword, s.valid)))
		user.GET("/:history_id", s.erh.HandleError(s.sess.Session(h.getHistory)))
	}

	company := rg.Group("/company")
	{
		company.GET("", s.erh.HandleError(s.sess.Session(h.getMyCompany)))
		company.PATCH("", s.erh.HandleError(session.HandleBody(h.updateCompany, s.sess.SessionFunc, s.valid)))
	}

	if s.mailSet {
		email := rg.Group("/email")
		{
			email.POST("/send-code", s.erh.HandleError(bind.HandleBody(h.sendCodeToEmail, s.valid)))
		}
	}
}

func (s *Setter) cors(c *gin.Context) {
	if origin := c.GetHeader("Origin"); origin != "" {
		c.Header("Access-Control-Allow-Origin", "http://localhost:3000")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		c.Header("Access-Control-Allow-Headers", "Content-Type, User-Agent, Accept-Language, Accept, Cache-Control, Content-Length, DomainName, Accept-Encoding, Connection, Set-Cookie, Cookie")
		c.Header("Access-Control-Expose-Headers", "Authorization")
		c.Header("Access-Control-Allow-Credentials", "true")
	} else {
		c.Header("Access-Control-Allow-Headers", c.GetHeader("Origin"))
	}
	return
}
