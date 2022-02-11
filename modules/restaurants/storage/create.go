package storage

import (
	"context"
	"locnguyen/common"
	"locnguyen/modules/restaurants/model"
)

func (s *sqlStore) Create(ctx context.Context, data *model.RestaurantCreate) error {
	db := s.db

	if err := db.Create(data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
