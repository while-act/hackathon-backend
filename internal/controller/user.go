package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/while-act/hackathon-backend/internal/controller/dao"
	"github.com/while-act/hackathon-backend/internal/controller/dto"
	"github.com/while-act/hackathon-backend/pkg/middleware/errs"
	"net/http"
	"strconv"
)

// GetMe godoc
// @Summary Get detail data about the user by session
// @Security ApiKeyAuth
// @Description Returns detail information about me
// @Tags Session
// @Success 200 {object} dao.Me "Info about session"
// @Failure 400 {object} errs.MyError "Validation error"
// @Failure 401 {object} errs.MyError "User isn't logged in"
// @Failure 500 {object} errs.MyError
// @Router /auth/session [get]
func (h *Handler) getMe(c *gin.Context, info *dao.Session) error {

	user, err := h.user.FindUserByID(info.ID)
	if err != nil {
		return err
	}

	c.JSON(http.StatusOK, user)
	return nil
}

// UpdateMe godoc
// @Summary Update user's data
// @Security ApiKeyAuth
// @Description Updates user's info
// @Tags User
// @Param updFields body dto.UpdateUser true "Fields to update"
// @Success 200 "Successfully Updated"
// @Failure 400 {object} errs.MyError "Validation error"
// @Failure 401 {object} errs.MyError "User isn't logged in"
// @Failure 500 {object} errs.MyError
// @Router /user [patch]
func (h *Handler) updateMe(c *gin.Context, updFields dto.UpdateUser, info *dao.Session) error {

	if err := h.user.UpdateUser(&updFields, info.ID); err != nil {
		return err
	}

	c.Status(http.StatusOK)
	return nil
}

// UpdatePassword godoc
// @Summary Update user's password
// @Description Updates user's password
// @Tags User
// @Param updPassword body dto.UpdatePassword true "Email with new password"
// @Success 200 "Successfully Updated"
// @Failure 400 {object} errs.MyError "Validation error"
// @Failure 500 {object} errs.MyError
// @Router /user/password [patch]
func (h *Handler) updatePassword(c *gin.Context, updPassword dto.UpdatePassword) error {

	if ok, err := h.auth.EqualsPopCode(updPassword.Email, updPassword.Code); err != nil {
		return err
	} else if !ok {
		return errs.CodeError.AddErr(err)
	}

	if err := h.user.UpdatePassword([]byte(updPassword.NewPassword), updPassword.Email); err != nil {
		return err
	}

	c.Status(http.StatusOK)
	return nil
}

// UpdateEmail godoc
// @Summary Update user's email
// @Security ApiKeyAuth
// @Description Updates user's email
// @Tags User
// @Param updEmail body dto.UpdateEmail true "New email with password"
// @Success 200 "Successfully Updated"
// @Failure 400 {object} errs.MyError "Validation error"
// @Failure 401 {object} errs.MyError "User isn't logged in"
// @Failure 500 {object} errs.MyError
// @Router /user/email [patch]
func (h *Handler) updateEmail(c *gin.Context, updEmail dto.UpdateEmail, info *dao.Session) error {

	if err := h.user.UpdateEmail([]byte(updEmail.Password), updEmail.NewEmail, info.ID); err != nil {
		return err
	}

	c.Status(http.StatusOK)
	return nil
}

// GetHistories godoc
// @Summary Get all user's histories
// @Description Returns array of user histories
// @Security ApiKeyAuth
// @Tags User
// @Success 200 {array} dao.Histories "History info"
// @Failure 401 {object} errs.MyError "User isn't logged in"
// @Failure 500 {object} errs.MyError
// @Router /user [get]
func (h *Handler) getHistories(c *gin.Context, info *dao.Session) error {

	history, err := h.user.GetAllHistory(info.ID)
	if err != nil {
		return err
	}
	c.JSON(http.StatusOK, history)
	return nil
}

// GetHistory godoc
// @Summary Generate PDF file from user history
// @Description Returns PDF file got from user history
// @Security ApiKeyAuth
// @Tags User
// @Param history_id path string true "Unique id from history"
// @Success 200 "PDF file"
// @Failure 400 {object} errs.MyError "Validation error"
// @Failure 401 {object} errs.MyError "User isn't logged in"
// @Failure 500 {object} errs.MyError
// @Router /user/{history_id} [get]
func (h *Handler) getHistory(c *gin.Context, info *dao.Session) error {

	historyIdStr := c.Param("history_id")

	err := binding.Validator.ValidateStruct(dto.HistoryId{Id: historyIdStr})
	if err != nil {
		return err
	}
	historyId, _ := strconv.Atoi(historyIdStr)

	history, err := h.user.GetOneHistory(historyId, info.ID)
	if err != nil {
		return err
	}
	dist, err := h.district.GetDistrict(history.DistrictTitle)
	if err != nil {
		return err
	}
	var tax float64
	if history.AccountingSupport {
		tax, err = h.tax.GetTax(&history.TaxationSystemOperations, &history.OperationType)
		if err != nil {
			return err
		}
	}

	err = h.pdf.GeneratePDF(c.Writer, h.pdf.CalcDB(history, dist, tax))
	if err != nil {
		return errs.PDFError.AddErr(err)
	}

	c.Header("Content-Disposition", "attachment; filename=result.pdf")
	c.Header("Content-Type", "application/pdf")
	c.Status(http.StatusOK)
	return nil
}
