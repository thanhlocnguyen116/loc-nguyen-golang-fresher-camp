package model

import (
	"errors"
	"locnguyen/common"
	"strings"
)

const EntityName = "Restaurant"

// CREATE TABLE `restaurants` (
// 	`id` int(11) NOT NULL,
// 	`name` varchar(255) NOT NULL,
// 	`address` varchar(255) NOT NULL,
// 	`cat_id` bigint(20) DEFAULT NULL,
// 	`ship_time` float DEFAULT '0',
// 	`free_ship` tinyint(1) DEFAULT '0',
// 	`has_liked` tinyint(1) DEFAULT '0',
// 	`rating` float DEFAULT NULL,
// 	`rating_count` int(11) DEFAULT NULL,
// 	`created_at` timestamp NOT NULL,
// 	`updated_at` timestamp NOT NULL,
// 	`status` tinyint(1) NOT NULL DEFAULT '1',
// 	PRIMARY KEY (`id`)
//   ) ENGINE=InnoDB DEFAULT CHARSET=utf8

type Restaurant struct {
	Id              string `json:"id,omitempty" gorm:"column:id"`
	Name            string `json:"name" gorm:"column:name"`
	Address         string `json:"address" gorm:"column:address"`
	common.SQLModel `json:",inline"`
}

func (Restaurant) TableName() string {
	return "restaurants"
}

type RestaurantUpdate struct {
	Name    *string `json:"name" gorm:"column:name"`
	Address *string `json:"address" gorm:"column:address"`
}

func (RestaurantUpdate) TableName() string {
	return Restaurant{}.TableName()
}

type RestaurantCreate struct {
	Id      string `json:"id,omitempty" gorm:"column:id"`
	Name    string `json:"name" gorm:"column:name"`
	Address string `json:"address" gorm:"column:address"`
}

func (RestaurantCreate) TableName() string {
	return Restaurant{}.TableName()
}

func (res *RestaurantCreate) Validate() error {
	res.Name = strings.TrimSpace(res.Name)

	if len(res.Name) == 0 {
		return errors.New("NO empty")
	}

	return nil

}
