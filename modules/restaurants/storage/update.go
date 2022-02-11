package storage

import (
	"context"
	"locnguyen/common"
	"locnguyen/modules/restaurants/model"
)

func (s *sqlStore) UpdateData(
	ctx context.Context,
	id int,
	data *model.RestaurantUpdate,
) error {
	db := s.db

	if err := db.Where("id = ?", id).Updates(data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
