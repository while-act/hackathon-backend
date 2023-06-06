package query

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/while-act/hackathon-backend/ent"
	"github.com/while-act/hackathon-backend/internal/controller/dao"
	"github.com/while-act/hackathon-backend/pkg/log"
	"strings"
	"time"
)

const (
	red    string = "\033[41m"
	green  string = "\033[42m"
	yellow string = "\033[43m"
	blue   string = "\033[44m"
	cyan   string = "\033[46m"
	white  string = "\033[47m"
)

type RespHandler struct {
	log *log.Logger
}

func NewQueryHandler() *RespHandler {
	return &RespHandler{log: log.NewLogger(log.InfoLevel, &log.TextFormatter{}, false)}
}

func (q *RespHandler) HandleQueries() gin.HandlerFunc {
	return func(c *gin.Context) {
		s := time.Now()
		c.Next()
		l := time.Since(s)

		defaultStatus, status := c.Writer.Status(), ""
		switch defaultStatus / 100 {
		case 2:
			status = setColor(defaultStatus, green)
		case 3:
			status = setColor(defaultStatus, blue)
		case 4:
			status = setColor(defaultStatus, yellow)
		case 5:
			status = setColor(defaultStatus, red)
		default:
			status = setColor(defaultStatus, white)
		}

		method := c.Request.Method
		switch method {
		case "POST":
			method = setColor(method, green)
		case "GET":
			method = setColor(method, blue)
		case "PATCH":
			method = setColor(method, cyan)
		case "DELETE":
			method = setColor(method, red)
		default:
			method = setColor(method, white)
		}

		q.log.Infof("%s| %12v |%6s| %s\n", status, l, method, c.Request.RequestURI)
	}
}

func setColor(text any, color string) string {
	return fmt.Sprintf("%s %v \033[0m", color, text)
}

func TransformIndustry(i *ent.Industry) *dao.Industry {
	return &dao.Industry{
		AvgWorkersNum:    i.AvgWorkersNum,
		AvgWorkersNumCad: i.AvgWorkersNumCad,
		AvgSalary:        i.AvgSalary,
		AvgSalaryCad:     i.AvgSalaryCad,
	}
}

func TransformToMe(user *ent.User) *dao.Me {
	return &dao.Me{
		Email:      user.Email[:1] + "**" + user.Email[strings.Index(user.Email, "@")-1:],
		Name:       user.Name,
		FirstName:  user.FirstName,
		LastName:   user.LastName,
		Role:       user.Role,
		FatherName: user.FatherName,
		Position:   user.Position,
		Country:    user.Country,
		City:       user.City,
		Biography:  user.Biography,
	}
}

func TransformCompany(c *ent.Company) *dao.Company {
	return &dao.Company{INN: c.Inn, Name: c.Name, Website: c.Website}
}
