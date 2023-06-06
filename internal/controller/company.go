package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/while-act/hackathon-backend/internal/controller/dao"
	"github.com/while-act/hackathon-backend/internal/controller/dto"
	"net/http"
)

// GetMyCompany godoc
// @Summary Get data about company by session
// @Security ApiKeyAuth
// @Description Returns information about company by session
// @Tags Company
// @Success 200 {object} dao.Company "Info about company"
// @Failure 401 {object} errs.MyError "User isn't logged in"
// @Failure 500 {object} errs.MyError
// @Router /company [get]
func (h *Handler) getMyCompany(c *gin.Context, info *dao.Session) error {
	user, err := h.user.FindUserByID(info.ID)
	if err != nil {
		return err
	}

	company, err := h.company.GetCompanyDTO(user.CompanyID)
	if err != nil {
		return err
	}

	c.JSON(http.StatusOK, company)
	return nil
}

// UpdateCompany godoc
// @Summary Update data about company
// @Description Updates information about company by INN
// @Tags Company
// @Security ApiKeyAuth
// @Param updCompany body dto.UpdateCompany true "Company"
// @Success 200 "OK"
// @Failure 400 {object} errs.MyError "Validation error"
// @Failure 401 {object} errs.MyError "User isn't logged in"
// @Failure 500 {object} errs.MyError
// @Router /company [patch]
func (h *Handler) updateCompany(c *gin.Context, updCompany dto.UpdateCompany, info *dao.Session) error {
	user, err := h.user.FindUserByID(info.ID)
	if err != nil {
		return err
	}

	if err = h.company.UpdateCompany(&updCompany, user.CompanyID); err != nil {
		return err
	}

	c.Status(http.StatusOK)
	return nil
}
