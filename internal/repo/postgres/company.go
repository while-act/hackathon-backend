package postgres

import (
	"context"
	"github.com/while-act/hackathon-backend/ent"
	"github.com/while-act/hackathon-backend/ent/company"
	"github.com/while-act/hackathon-backend/internal/controller/dao"
	"github.com/while-act/hackathon-backend/internal/controller/dto"
	"github.com/while-act/hackathon-backend/pkg/middleware/query"
)

type CompanyStorage struct {
	companyClient *ent.CompanyClient
}

func NewCompanyStorage(companyClient *ent.CompanyClient) *CompanyStorage {
	return &CompanyStorage{companyClient: companyClient}
}

func (r *CompanyStorage) CreateCompany(ctx context.Context, inn string, name, website *string) (*ent.Company, error) {
	return r.companyClient.Create().
		SetInn(inn).SetNillableName(name).
		SetNillableWebsite(website).Save(ctx)
}

func (r *CompanyStorage) GetCompanyDTO(ctx context.Context, id int) (*dao.Company, error) {
	comp, err := r.companyClient.Query().Where(company.ID(id)).Select(
		company.FieldInn, company.FieldWebsite, company.FieldName).Only(ctx)

	if err == nil {
		return query.TransformCompany(comp), nil
	}

	return nil, err
}

func (r *CompanyStorage) GetCompany(ctx context.Context, id int) (*ent.Company, error) {
	return r.companyClient.Get(ctx, id)
}

func (r *CompanyStorage) UpdateCompany(ctx context.Context, updateCompany *dto.UpdateCompany, id int) error {
	return r.companyClient.UpdateOneID(id).
		SetNillableName(updateCompany.Name).
		SetNillableWebsite(updateCompany.Website).Exec(ctx)
}
