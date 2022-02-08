package usermodel

import (
	"errors"
	"locnguyen/common"
	"strings"
)

const EntityName = "User"

type User struct {
	common.SQLModel `json:",inline"`
	User_id         int    `json:"user_id,omitempty" gorm:"column:user_id"`
	Name            string `json:"name" gorm:"column:name"`
	Password        string `json:"password" gorm:"column:password"`
	Email           string `json:"email" gorm:"column:email"`
	Phone           int    `json:"phone" gorm:"column:phone"`
	Address         string `json:"address" gorm:"column:address"`
}

func (User) TableName() string {
	return "users"
}

type UserUpdate struct {
	Name *string `json:"name" gorm:"column:name;"`
}

func (UserUpdate) TableName() string {
	return User{}.TableName()
}

type UserCreate struct {
	User_id  int    `json:"user_id,omitempty" gorm:"column:user_id"`
	Name     string `json:"name" gorm:"column:name"`
	Password string `json:"password" gorm:"column:password"`
	Email    string `json:"email" gorm:"column:email"`
	Phone    int    `json:"phone" gorm:"column:phone"`
	Address  string `json:"address" gorm:"column:address"`
}

func (UserCreate) TableName() string {
	return User{}.TableName()
}

func (res *UserCreate) Validate() error {
	res.Name = strings.TrimSpace(res.Name)

	if len(res.Name) == 0 {
		return errors.New("sao lại để trống?")
	}

	return nil
}
