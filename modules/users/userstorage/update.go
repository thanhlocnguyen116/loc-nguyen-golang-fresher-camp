package userstorage

import (
	"context"
	"locnguyen/common"
	"locnguyen/modules/users/usermodel"
)

func (s *sqlStore) UpdateData(
	ctx context.Context,
	user_id int,
	data *usermodel.UserUpdate,
) error {
	db := s.db

	if err := db.Where("user_id = ?", user_id).Updates(data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
