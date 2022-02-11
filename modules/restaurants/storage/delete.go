package storage

import (
	"context"
	"locnguyen/common"
	"locnguyen/modules/restaurants/model"
)

func (s *sqlStore) SoftDeleteData(
	ctx context.Context,
	id int,
) error {
	db := s.db

	if err := db.Table(model.Restaurant{}.TableName()).Where("id = ?", id).Updates(map[string]interface{}{
		"status": 0,
	}).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
