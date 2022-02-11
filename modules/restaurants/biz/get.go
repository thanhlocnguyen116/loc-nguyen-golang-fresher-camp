package biz

import (
	"context"
	"locnguyen/common"
	"locnguyen/modules/restaurants/model"
	"locnguyen/modules/users/usermodel"
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
			return nil, common.ErrCannotGetEntity(usermodel.EntityName, err)
		}

		return nil, common.ErrCannotGetEntity(usermodel.EntityName, err)
	}

	if data.Status == 0 {
		return nil, common.ErrEntityDeleted(usermodel.EntityName, nil)
	}

	return data, err
}
