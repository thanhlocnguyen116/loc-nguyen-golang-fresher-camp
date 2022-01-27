package usersbiz

import (
	"context"
	"locnguyen/modules/users/usermodel"
)

type GetUserStore interface {
	FindDataByCondition(
		ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string,
	) (*usermodel.User, error)
}

type getUserBiz struct {
	store GetUserStore
}

func NewGetUserBiz(store GetUserStore) *getUserBiz {
	return &getUserBiz{store: store}
}

func (biz *getUserBiz) GetUser(ctx context.Context, id int) (*usermodel.User, error) {
	data, err := biz.store.FindDataByCondition(ctx, map[string]interface{}{"id": id})

	return data, err
}
