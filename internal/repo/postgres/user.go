package postgres

import (
	"context"
	"github.com/wtkeqrf0/while.act/ent"
)

type UserStorage struct {
	userClient *ent.UserClient
}

func NewUserStorage(userClient *ent.UserClient) *UserStorage {
	return &UserStorage{userClient: userClient}
}

func (r *UserStorage) FindUserByID(ctx context.Context, id int) (*ent.User, error) {
	return r.userClient.Get(ctx, id)
}
