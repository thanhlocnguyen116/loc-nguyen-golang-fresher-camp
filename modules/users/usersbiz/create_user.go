package usersbiz

import (
	"context"
	"locnguyen/modules/users/usermodel"
)

type CreateUserStore interface {
	Create(ctx context.Context, data *usermodel.UserCreate) error
}

type createUserBiz struct {
	store CreateUserStore
}

func NewCreateUserBiz(store CreateUserStore) *createUserBiz {
	return &createUserBiz{store: store}
}

func (biz *createUserBiz) CreateUser(ctx context.Context, data *usermodel.UserCreate) error {
	if err := data.Validate(); err != nil {
		return err
	}

	err := biz.store.Create(ctx, data)

	return err
}
