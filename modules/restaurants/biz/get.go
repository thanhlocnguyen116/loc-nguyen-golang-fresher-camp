package biz

import (
	"context"
	"locnguyen/common"
	"locnguyen/modules/restaurants/model"
)

type GetRestaurantStore interface {
	FindDataByCondition(
		ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string,
	) (*model.Restaurant, error)
}

type getRestaurantBiz struct {
	store GetRestaurantStore
}

func NewGetRestaurantBiz(store GetRestaurantStore) *getRestaurantBiz {
	return &getRestaurantBiz{store: store}
}

func (biz *getRestaurantBiz) GetRestaurant(ctx context.Context, id int) (*model.Restaurant, error) {
	data, err := biz.store.FindDataByCondition(ctx, map[string]interface{}{"id": id})

	if err != nil {
		if err != common.RecordNotFound {
			return nil, common.ErrCannotGetEntity(model.EntityName, err)
		}

		return nil, common.ErrCannotGetEntity(model.EntityName, err)
	}

	if data.Status == 0 {
		return nil, common.ErrEntityDeleted(model.EntityName, nil)
	}

	return data, err
}
