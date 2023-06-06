package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/while-act/hackathon-backend/internal/controller/dao"
	"github.com/while-act/hackathon-backend/internal/controller/dto"
	"github.com/while-act/hackathon-backend/pkg/middleware/errs"
	"net/http"
)

// SaveCalcData godoc
// @Summary Save calc data to history
// @Security ApiKeyAuth
// @Description Saves given values to user's history
// @Tags Calc
// @Param from body dto.History true "Completed application form"
// @Success 200 "OK"
// @Failure 400 {object} errs.MyError "Validation error"
// @Failure 401 {object} errs.MyError "User isn't logged in"
// @Failure 500 {object} errs.MyError
// @Router /calc/save [post]
func (h *Handler) saveCalcData(c *gin.Context, history dto.History, info *dao.Session) error {

	id, err := h.business.GetBusiness(history.BusinessActivity)
	if err != nil {
		return err
	}

	if err = h.history.CreateHistory(&history, id, info.ID); err != nil {
		return err
	}

	c.Status(http.StatusOK)
	return nil
}

// CalcData godoc
// @Summary Generate PDF from body
// @Description Returns PDF file, gotten from body
// @Tags Calc
// @Param from body dto.History true "Completed application form"
// @Success 200 "PDF file"
// @Failure 400 {object} errs.MyError "Validation error"
// @Failure 500 {object} errs.MyError
// @Router /calc [post]
func (h *Handler) calcData(c *gin.Context, history dto.History) error {

	dist, err := h.district.GetDistrict(history.DistrictTitle)
	if err != nil {
		return err
	}
	var tax float64
	if history.AccountingSupport {
		tax, err = h.tax.GetTax(history.TaxationSystemOperations, history.OperationsType)
		if err != nil {
			return err
		}
	}

	err = h.pdf.GeneratePDF(c.Writer, h.pdf.CalcDTO(&history, dist, tax))
	if err != nil {
		return errs.PDFError.AddErr(err)
	}

	c.Header("Content-Disposition", "attachment; filename=result.pdf")
	c.Header("Content-Type", "application/pdf")
	c.Status(http.StatusOK)
	return nil
}

// GetIndustryInfo godoc
// @Summary Get data about industry
// @Description Returns detail information about industry
// @Tags Calc
// @Param industry path string true "Industry Branch"
// @Success 200 {object} dao.Industry "Info about industry"
// @Failure 400 {object} errs.MyError "Validation error"
// @Failure 500 {object} errs.MyError
// @Router /calc/{industry} [get]
func (h *Handler) getIndustryInfo(c *gin.Context, ind string) error {
	industry, err := h.industry.GetIndustry(ind)
	if err != nil {
		return err
	}

	c.JSON(http.StatusOK, industry)
	return nil
}
