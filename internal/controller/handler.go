package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/while-act/hackathon-backend/ent"
	"github.com/while-act/hackathon-backend/internal/controller/dao"
	"github.com/while-act/hackathon-backend/internal/controller/dto"
	"github.com/while-act/hackathon-backend/internal/service"
	"github.com/while-act/hackathon-backend/pkg/conf"
	"io"
	"time"
)

var cfg = conf.GetConfig()

// UserService interacts with the users table
type UserService interface {
	FindUserByID(id int) (*dao.Me, error)

	UpdateUser(updateUser *dto.UpdateUser, id int) error
	UpdatePassword(newPassword []byte, email string) error
	UpdateEmail(password []byte, newEmail string, id int) error

	GetAllHistory(userId int) ([]*dao.Histories, error)
	GetOneHistory(historyId int, userId int) (*ent.History, error)

	ContainsKeys(keys ...string) (int64, error)
	SetVariable(key string, value any, exp time.Duration) error
}

type CompanyService interface {
	CreateCompany(inn string, name, website *string) (*ent.Company, error)
	GetCompany(id int) (*ent.Company, error)
	GetCompanyDTO(id int) (*dao.Company, error)
	UpdateCompany(updateCompany *dto.UpdateCompany, id int) error
}

type AuthService interface {
	CreateUserWithPassword(auth *dto.SignUp, company *ent.Company) (*ent.User, error)
	AuthUserByEmail(email string) (*ent.User, error)
	DelKeys(keys ...string)

	EqualsPopCode(email string, code string) (bool, error)
	SetCodes(key string, value ...any) error
}

type IndustryService interface {
	GetIndustry(title string) (*dao.Industry, error)
}

type DistrictService interface {
	GetDistrict(title string) (*ent.District, error)
}

type HistoryService interface {
	GetHistory(historyId int) (*ent.History, error)
	CreateHistory(h *dto.History, busactId *int, id int) error
}

type BusinessActivityService interface {
	GetBusiness(bus *dto.BusinessActivity) (*int, error)
}

type TaxService interface {
	GetTax(num *int, tax *string) (float64, error)
}

type PDFGenerator interface {
	CalcDTO(h *dto.History, dist *ent.District, tax float64) service.Params
	CalcDB(h *ent.History, dist *ent.District, tax float64) service.Params
	GeneratePDF(out io.Writer, data service.Params) error
}

type Sessions interface {
	SetNewCookie(id int, c *gin.Context)
	GenerateSecretCode() string
}

type MailSender interface {
	SendEmail(subj, body, from string, to ...string) error
}

type Handler struct {
	user     UserService
	company  CompanyService
	auth     AuthService
	history  HistoryService
	industry IndustryService
	district DistrictService
	tax      TaxService
	business BusinessActivityService
	pdf      PDFGenerator
	session  Sessions
	mail     MailSender
}

func NewHandler(user UserService, company CompanyService, auth AuthService, history HistoryService, industry IndustryService, district DistrictService, tax TaxService, business BusinessActivityService, pdf PDFGenerator, session Sessions, mail MailSender) *Handler {
	return &Handler{user: user, company: company, auth: auth, history: history, industry: industry, district: district, tax: tax, business: business, pdf: pdf, session: session, mail: mail}
}
