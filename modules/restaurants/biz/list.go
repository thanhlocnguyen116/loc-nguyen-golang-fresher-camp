package biz

import (
	"context"
	"locnguyen/common"
	"locnguyen/modules/restaurants/model"
)

type ListRestaurantStore interface {
	ListDataByCondition(ctx context.Context,
		conditions map[string]interface{},
		filter *model.Filter,
		paging *common.Paging,
		moreKeys ...string,
	) ([]model.Restaurant, error)
}

type listRestaurantBiz struct {
	store ListRestaurantStore
}

func NewListRestaurantBiz(store ListRestaurantStore) *listRestaurantBiz {
	return &listRestaurantBiz{store: store}
}

func (biz *listRestaurantBiz) ListRestaurant(
	ctx context.Context,
	filter *model.Filter,
	paging *common.Paging,
) ([]model.Restaurant, error) {
	result, err := biz.store.ListDataByCondition(ctx, nil, filter, paging)

	return result, err
}
