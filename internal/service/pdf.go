package service

import (
	"bytes"
	pdff "github.com/SebastiaanKlippert/go-wkhtmltopdf"
	"github.com/while-act/hackathon-backend/ent"
	"github.com/while-act/hackathon-backend/internal/controller/dto"
	"github.com/while-act/hackathon-backend/pkg/log"
	"html/template"
	"io"
)

type PDF struct {
	t *template.Template
}

type Params struct {
	Name              string
	IndustryBranch    string
	OrganizationType  string
	FullTimeEmployers int
	LandArea          float64
	WageFund          float64
	InsurancePayment  float64
	IncomeTax         float64
	LandValueMin      float64
	LandValueMax      float64
	LandValue         float64
	Equipment         float64
	Taxes             float64
	SocialInsurance   float64
	PatentCost        float64
	Other             string
	Total             float64
}

func NewPDF(templatePath string) *PDF {
	t, err := template.ParseFiles(templatePath)
	if err != nil {
		log.Fatal(err)
	}

	pdff.SetPath("configs/wkhtmltopdf.exe")

	return &PDF{t: t}
}

func (r *PDF) CalcDTO(h *dto.History, dist *ent.District, tax float64, patent float64) Params {
	p := Params{
		Name:              h.Name,
		Other:             h.Other,
		IndustryBranch:    h.IndustryBranch,
		OrganizationType:  h.OrganizationLegal,
		FullTimeEmployers: h.FullTimeEmployees,
		LandArea:          h.LandArea,
	}

	wageFund := float64(h.FullTimeEmployees) * h.AvgSalary * 12
	p.InsurancePayment = wageFund * 0.3
	p.IncomeTax = wageFund * 0.13
	p.WageFund = wageFund + p.InsurancePayment + p.IncomeTax
	p.SocialInsurance = wageFund * 30
	p.Total += p.WageFund + p.SocialInsurance + p.InsurancePayment + p.IncomeTax

	if h.IsBuy {
		p.LandValue = h.LandArea * dist.AvgCadastralVal
		p.LandValueMin = h.ConstructionFacilitiesArea * 80
		p.LandValueMax = h.ConstructionFacilitiesArea * 120
	} else {
		p.LandValue = h.LandArea * dist.AvgCadastralVal * 0.003
	}
	p.Total += p.LandValue
	for _, v := range h.Equipment {
		p.Equipment += v.Price
	}
	p.Total += p.Equipment
	if h.AccountingSupport {
		p.Taxes = tax + (0.5 * float64(h.FullTimeEmployees))
		p.Total += p.Taxes
	}

	if h.PatentCalc && patent != 0 {
		p.PatentCost = patent
		p.Total += patent
	}

	return p
}

func (r *PDF) CalcDB(h *ent.History, dist *ent.District, tax float64, patent float64) Params {
	p := Params{
		Name:              h.Name,
		Other:             h.Other,
		IndustryBranch:    h.IndustryBranch,
		OrganizationType:  h.OrganizationalLegal,
		FullTimeEmployers: h.FullTimeEmployees,
		LandArea:          h.LandArea,
	}

	wageFund := float64(h.FullTimeEmployees) * h.AvgSalary * 12
	p.InsurancePayment = wageFund * 0.3
	p.IncomeTax = wageFund * 0.13
	p.WageFund = wageFund + p.InsurancePayment + p.IncomeTax
	p.SocialInsurance = wageFund * 30
	p.Total += p.WageFund + p.SocialInsurance + p.InsurancePayment + p.IncomeTax

	if h.IsBuy {
		p.LandValue = h.LandArea * dist.AvgCadastralVal
		p.LandValueMin = h.ConstructionFacilitiesArea * 80
		p.LandValueMax = h.ConstructionFacilitiesArea * 120
	} else {
		p.LandValue = h.LandArea * dist.AvgCadastralVal * 0.003
	}
	p.Total += p.LandValue
	for _, v := range h.Equipment {
		p.Equipment += v.Price
	}
	p.Total += p.Equipment
	if h.AccountingSupport {
		p.Taxes = tax + (0.5 * float64(h.FullTimeEmployees))
		p.Total += p.Taxes
	}

	if h.PatentCalc && patent != 0 {
		p.PatentCost = patent
		p.Total += patent
	}

	return p
}

func (r *PDF) GeneratePDF(out io.Writer, data Params) error {

	pdf, err := pdff.NewPDFGenerator()
	if err != nil {
		return err
	}

	buf := new(bytes.Buffer)
	if err = r.t.Execute(buf, data); err != nil {
		return err
	}

	pdf.NoPdfCompression.Set(true)
	pdf.SetOutput(out)
	pdf.PageSize.Set(pdff.PageSizeA4)
	pdf.AddPage(pdff.NewPageReader(buf))

	return pdf.Create()
}
