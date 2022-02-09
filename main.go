package main

import (
	"locnguyen/component"
	"locnguyen/middleware"
	"locnguyen/modules/users/usertransport/ginuser"
	"log"
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

	// r.GET("/ping", func(c *gin.Context) {
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"message": "ponk",
	// 	})
	// })

	// CRUD

	users := r.Group("/users")
	{
		users.POST("", ginuser.CreateUser(appCtx))
		users.GET("/:id", ginuser.GetUser(appCtx))
		users.GET("", ginuser.ListUser(appCtx))
		users.PATCH("/:id", ginuser.UpdateUser(appCtx))
		users.DELETE("/:id", ginuser.DeleteUser(appCtx))
	}

	return r.Run()

}
