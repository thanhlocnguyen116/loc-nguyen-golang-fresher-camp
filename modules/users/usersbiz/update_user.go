package usersbiz

import (
	"context"
	"errors"
	"locnguyen/modules/users/usermodel"
)

type UpdateUserStore interface {
	FindDataByCondition(
		ctx context.Context,
		condition map[string]interface{},
		morekeys ...string,
	) (*usermodel.User, error)

	UpdateData(
		ctx context.Context,
		user_id int,
		data *usermodel.UserUpdate,
	) error
}

type updateUserBiz struct {
	store UpdateUserStore
}

func NewUpdateUserBiz(store UpdateUserStore) *updateUserBiz {
	return &updateUserBiz{store: store}
}

func (biz *updateUserBiz) UpdateUser(ctx context.Context, user_id int, data *usermodel.UserUpdate) error {
	oldData, err := biz.store.FindDataByCondition(ctx, map[string]interface{}{"user_id": user_id})

	if err != nil {
		return err
	}

	if oldData.Status == 0 {
		return errors.New("data deleted")
	}

	if err := biz.store.UpdateData(ctx, user_id, data); err != nil {
		return nil
	}

	return nil
}
