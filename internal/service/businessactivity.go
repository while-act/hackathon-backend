package service

import (
	"context"
	"github.com/while-act/hackathon-backend/ent"
)

type BusinessPostgres interface {
	GetBusiness(ctx context.Context, bus string) (*ent.BusinessActivity, error)
}

type BusinessService struct {
	postgres BusinessPostgres
}

func NewBusinessService(postgres BusinessPostgres) *BusinessService {
	return &BusinessService{postgres: postgres}
}

func (b *BusinessService) GetBusiness(bus string) (*ent.BusinessActivity, error) {
	return b.postgres.GetBusiness(context.Background(), bus)
}
