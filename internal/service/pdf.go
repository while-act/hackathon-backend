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
	t []*template.Template
}

type Params struct {
	Name              string
	IndustryBranch    string
	OrganizationType  string
	Other             string
	DistrictTitle     string
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
	Total             float64
	MaxTotal          float64
}

func NewPDF(templatePath string) *PDF {
	t1, _ := template.ParseFiles(templatePath + "template.html")
	t2, _ := template.ParseFiles(templatePath + "template2.html")
	t3, _ := template.ParseFiles(templatePath + "template3.html")
	t4, _ := template.ParseFiles(templatePath + "template4.html")
	t5, _ := template.ParseFiles(templatePath + "template5.html")

	t := []*template.Template{t1, t2, t3, t4, t5}
	t1, err := template.ParseFiles(templatePath+"template.html", templatePath+"template2.html",
		templatePath+"template3.html", templatePath+"template4.html", templatePath+"template5.html")
	if err != nil {
		log.Fatal(err)
	}

	pdff.SetPath("configs/wkhtmltopdf.exe")

	return &PDF{t: t}
}

func (r *PDF) CalcDTO(h *dto.History, dist *ent.District, tax float64, patent float64) Params {
	p := Params{
		DistrictTitle:     h.DistrictTitle,
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
	if h.AccountingSupport && tax != 0 {
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

	data.LandArea = data.LandArea / 1000
	data.WageFund = data.WageFund / 1000
	data.InsurancePayment = data.InsurancePayment / 1000
	data.IncomeTax = data.IncomeTax / 1000
	data.LandValueMin = data.LandValueMin / 1000
	data.LandValueMax = data.LandValueMax / 1000
	data.LandValue = data.LandValue / 1000
	data.Equipment = data.Equipment / 1000
	data.Taxes = data.Taxes / 1000
	data.SocialInsurance = data.SocialInsurance / 1000
	data.PatentCost = data.PatentCost / 1000
	data.Total = data.Total / 1000
	data.MaxTotal = data.Total * 1.3

	pdf.PageSize.Set(pdff.PageSizeA4)
	pdf.Dpi.Set(300)
	pdf.SetOutput(out)

	buf := new(bytes.Buffer)
	for _, v := range r.t {
		if err = v.Execute(buf, data); err != nil {
			return err
		}
	}
	page := pdff.NewPageReader(buf)
	page.HeaderSpacing.Set(10)
	page.Encoding.Set("UTF8")
	pdf.AddPage(pdff.NewPageReader(buf))

	return pdf.Create()
}
