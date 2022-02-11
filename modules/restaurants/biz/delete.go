package biz

import (
	"context"
	"errors"
	"locnguyen/modules/restaurants/model"
)

type DeleteRestaurantStore interface {
	FindDataByCondition(
		ctx context.Context,
		condition map[string]interface{},
		morekeys ...string,
	) (*model.Restaurant, error)

	SoftDeleteData(
		ctx context.Context,
		id int,
	) error
}

type deleteRestaurantBiz struct {
	store DeleteRestaurantStore
}

func NewDeleteRestaurantBiz(store DeleteRestaurantStore) *deleteRestaurantBiz {
	return &deleteRestaurantBiz{store: store}
}

func (biz *deleteRestaurantBiz) DeleteRestaurant(ctx context.Context, id int) error {
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
