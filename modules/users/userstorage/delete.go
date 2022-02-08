package userstorage

import (
	"context"
	"locnguyen/common"
	"locnguyen/modules/users/usermodel"
)

func (s *sqlStore) SoftDeleteData(
	ctx context.Context,
	user_id int,
) error {
	db := s.db

	if err := db.Table(usermodel.User{}.TableName()).Where("user_id = ?", user_id).Updates(map[string]interface{}{
		"status": 0,
	}).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
