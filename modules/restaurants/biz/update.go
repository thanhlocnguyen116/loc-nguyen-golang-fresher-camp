package biz

import (
	"context"
	"errors"
	"locnguyen/modules/restaurants/model"
)

type UpdateRestaurantStore interface {
	FindDataByCondition(
		ctx context.Context,
		condition map[string]interface{},
		morekeys ...string,
	) (*model.Restaurant, error)

	UpdateData(
		ctx context.Context,
		id int,
		data *model.RestaurantUpdate,
	) error
}

type updateRestaurantBiz struct {
	store UpdateRestaurantStore
}

func NewUpdateRestaurantBiz(store UpdateRestaurantStore) *updateRestaurantBiz {
	return &updateRestaurantBiz{store: store}
}

func (biz *updateRestaurantBiz) UpdateRestaurant(ctx context.Context, id int, data *model.RestaurantUpdate) error {
	oldData, err := biz.store.FindDataByCondition(ctx, map[string]interface{}{"id": id})

	if err != nil {
		return err
	}

	if oldData.Status == 0 {
		return errors.New("data deleted")
	}

	if err := biz.store.UpdateData(ctx, id, data); err != nil {
		return nil
	}

	return nil
}
