package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func init() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Không load được file .env")
	}
}

// CREATE TABLE `users` (
// 	`user_id` int(11) NOT NULL AUTO_INCREMENT,
// 	`name` varchar(100) NOT NULL,
// 	`password` varchar(100) NOT NULL,
// 	`email` varchar(255) NOT NULL,
// 	`phone` int(11) NOT NULL,
// 	`address` varchar(255) NOT NULL,
// 	`image` json DEFAULT NULL,
// 	`created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
// 	`updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
// 	PRIMARY KEY (`user_id`),
// 	UNIQUE KEY `id` (`user_id`)
//   ) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8

type User struct {
	User_id  int    `json:"user_id,omitempty" gorm:"column:user_id"`
	Name     string `json:"name" gorm:"column:name"`
	Password string `json:"password" gorm:"column:password"`
	Email    string `json:"email" gorm:"column:email"`
	Phone    int    `json:"phone" gorm:"column:phone"`
	Address  string `json:"address" gorm:"column:address"`
}

func main() {

	dsn := os.Getenv("DBconnectionString")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	fmt.Println(db, err)
}
