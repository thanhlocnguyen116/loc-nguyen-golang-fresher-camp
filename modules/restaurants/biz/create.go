package biz

import (
	"context"
	"locnguyen/modules/restaurants/model"
)

type CreateRestaurantStore interface {
	Create(ctx context.Context, data *model.RestaurantCreate) error
}

type CreateRestaurantBiz struct {
	store CreateRestaurantStore
}

func NewCreateRestaurantBiz(store CreateRestaurantStore) *CreateRestaurantBiz {
	return &CreateRestaurantBiz{store: store}
}

func (biz *CreateRestaurantBiz) CreateRestaurant(ctx context.Context, data *model.RestaurantCreate) error {
	if err := data.Validate(); err != nil {
		return err
	}

	err := biz.store.Create(ctx, data)

	return err
}
