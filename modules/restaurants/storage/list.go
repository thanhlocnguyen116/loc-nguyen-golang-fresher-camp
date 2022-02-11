package storage

import (
	"context"
	"locnguyen/common"
	"locnguyen/modules/restaurants/model"
)

func (s *sqlStore) ListDataByCondition(ctx context.Context,
	conditions map[string]interface{},
	filter *model.Filter,
	paging *common.Paging,
	moreKeys ...string,
) ([]model.Restaurant, error) {
	var result []model.Restaurant

	db := s.db

	for i := range moreKeys {
		db = db.Preload(moreKeys[i])
	}

	db = db.Table(model.Restaurant{}.TableName()).Where(conditions).Where("status in (1)")

	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	if err := db.
		Offset((paging.Page - 1) * paging.Limit).
		Limit(paging.Limit).
		Find(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return result, nil
}
