package main

import (
	"locnguyen/component"
	"locnguyen/middleware"
	"locnguyen/modules/users/usertransport/ginuser"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
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

type UserUpdate struct {
	Name *string `json:"name" gorm:"column:name"`
}

func (User) TableName() string {
	return "users"
}

func (UserUpdate) TableName() string {
	return User{}.TableName()
}

func main() {
	dsn := os.Getenv("DBconnectionString")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	if err := runService(db); err != nil {
		log.Fatalln(err)
	}

}

func runService(db *gorm.DB) error {

	appCtx := component.NewAppContext(db)
	r := gin.Default()

	r.Use(middleware.Recover(appCtx))

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "ponk",
		})
	})

	// CRUD

	users := r.Group("/users")
	{
		users.POST("", ginuser.CreateUser(appCtx))
		users.GET("/:user_id", ginuser.GetUser(appCtx))
		users.GET("", ginuser.ListUser(appCtx))
		users.PATCH("/:user_id", ginuser.UpdateUser(appCtx))
		users.DELETE("/:user_id", ginuser.DeleteUser(appCtx))
	}

	return r.Run()
}
