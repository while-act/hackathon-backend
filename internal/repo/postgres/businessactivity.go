package postgres

import (
	"context"
	"github.com/while-act/hackathon-backend/ent"
)

type BusinessStorage struct {
	businessClient *ent.BusinessActivityClient
}

func NewBusinessStorage(businessClient *ent.BusinessActivityClient) *BusinessStorage {
	return &BusinessStorage{businessClient: businessClient}
}

func (b *BusinessStorage) GetBusiness(ctx context.Context, bus string) (*ent.BusinessActivity, error) {
	return b.businessClient.Get(ctx, bus)
}
