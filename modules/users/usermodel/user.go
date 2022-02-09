package usermodel

import (
	"errors"
	"locnguyen/common"
	"strings"
)

const EntityName = "User"

type User struct {
	common.SQLModel `json:",inline"`
	Id              int    `json:"id,omitempty" gorm:"column:id"`
	Email           string `json:"email" gorm:"column:email"`
	Password        string `json:"password" gorm:"column:password"`
	FirstName       string `json:"first_name" gorm:"column:first_name"`
	LastName        string `json:"last_name" gorm:"column:last_name"`
	Role            string `json:"role" gorm:"column:role"`
}

func (User) TableName() string {
	return "users"
}

type UserUpdate struct {
	FirstName *string `json:"first_name" gorm:"column:first_name"`
	LastName  *string `json:"last_name" gorm:"column:last_name"`
}

func (UserUpdate) TableName() string {
	return User{}.TableName()
}

type UserCreate struct {
	Id        int    `json:"id,omitempty" gorm:"column:id"`
	Email     string `json:"email" gorm:"column:email"`
	Password  string `json:"password" gorm:"column:password"`
	FirstName string `json:"first_name" gorm:"column:first_name"`
	LastName  string `json:"last_name" gorm:"column:last_name"`
	Role      string `json:"role" gorm:"column:role"`
}

func (UserCreate) TableName() string {
	return User{}.TableName()
}

func (res *UserCreate) Validate() error {
	res.Email = strings.TrimSpace(res.Email)

	if len(res.Email) == 0 {
		return errors.New("sao lại để trống?")
	}

	return nil
}
