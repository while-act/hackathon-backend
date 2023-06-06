package dto

type History struct {
	Name                       string            `json:"name,omitempty" validate:"required,gte=2,lte=200" example:"ООО 'Парк'"`
	OrganizationLegal          string            `json:"organizationLegal,omitempty" validate:"required,enum=ИП*ООО" example:"ИП"`
	IndustryBranch             string            `json:"industryBranch,omitempty" validate:"required" example:"Авиационная_промышленность"`
	FullTimeEmployees          int               `json:"fullTimeEmployees,omitempty" validate:"required,gte=0" example:"50"`
	AvgSalary                  float64           `json:"avgSalary,omitempty" validate:"required,gte=0" example:"3058.12"`
	DistrictTitle              string            `json:"districtTitle,omitempty" validate:"required" example:"ЗАО"`
	LandArea                   float64           `json:"landArea,omitempty" validate:"required,gte=0" example:"120"`
	IsBuy                      bool              `json:"isBuy,omitempty" example:"true"`
	ConstructionFacilitiesArea float64           `json:"constructionFacilitiesArea,omitempty" validate:"required,gte=0" example:"50"`
	BuildingTypes              []string          `json:"buildingTypes,omitempty" example:"Энергетические"`
	Equipment                  []Equipment       `json:"equipment,omitempty" validate:"omitempty"`
	AccountingSupport          bool              `json:"accountingSupport,omitempty" example:"true"`
	TaxationSystemOperations   *int              `json:"taxationSystemOperations,omitempty" validate:"omitempty"`
	OperationsType             *string           `json:"operationsType,omitempty" validate:"omitempty,enum=usn6*usn15*osn" example:"usn6"`
	PatentCalc                 bool              `json:"patentCalc,omitempty" example:"true"`
	BusinessActivity           *BusinessActivity `json:"business_activity,omitempty"`
	Other                      string            `json:"other,omitempty" example:"I want some cookies"`
}

type HistoryId struct {
	Id string `json:"id,omitempty" validate:"required,number" example:"12"`
}
