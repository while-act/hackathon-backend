package postgres

import (
	"context"
	"github.com/while-act/hackathon-backend/ent"
	"github.com/while-act/hackathon-backend/ent/industry"
	"github.com/while-act/hackathon-backend/internal/controller/dao"
	"github.com/while-act/hackathon-backend/pkg/middleware/query"
)

type IndustryStorage struct {
	industryClient *ent.IndustryClient
}

func NewIndustryStorage(industryClient *ent.IndustryClient) *IndustryStorage {
	return &IndustryStorage{industryClient: industryClient}
}

func (i *IndustryStorage) GetIndustry(ctx context.Context, title string) (*dao.Industry, error) {
	ind, err := i.industryClient.Query().Where(industry.ID(title)).Select(
		industry.FieldAvgSalary, industry.FieldAvgSalaryCad,
		industry.FieldAvgWorkersNumCad, industry.FieldAvgWorkersNum,
	).Only(ctx)
	if err == nil {
		return query.TransformIndustry(ind), nil
	}
	return nil, err
}
