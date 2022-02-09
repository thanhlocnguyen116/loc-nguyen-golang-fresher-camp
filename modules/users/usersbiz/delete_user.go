package usersbiz

import (
	"context"
	"errors"
	"locnguyen/modules/users/usermodel"
)

type DeleteUserStore interface {
	FindDataByCondition(
		ctx context.Context,
		condition map[string]interface{},
		morekeys ...string,
	) (*usermodel.User, error)

	SoftDeleteData(
		ctx context.Context,
		id int,
	) error
}

type deleteUserBiz struct {
	store DeleteUserStore
}

func NewDeleteUserBiz(store DeleteUserStore) *deleteUserBiz {
	return &deleteUserBiz{store: store}
}

func (biz *deleteUserBiz) DeleteUser(ctx context.Context, id int) error {
	oldData, err := biz.store.FindDataByCondition(ctx, map[string]interface{}{"id": id})

	if err != nil {
		return err
	}

	if oldData.Status == 0 {
		return errors.New("data deleted")
	}

	if err := biz.store.SoftDeleteData(ctx, id); err != nil {
		return nil
	}

	return nil
}
