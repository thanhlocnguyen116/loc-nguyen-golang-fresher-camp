package main

import (
	"locnguyen/component"
	"locnguyen/component/uploadprovider"
	"locnguyen/middleware"
	"locnguyen/modules/restaurants/transport/ginrestaurant"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//set environment file
func init() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Không load được file .env")
	}
}

func main() {

	//connect DB
	dsn := os.Getenv("DBconnectionString")

	//S3
	s3BucketName := os.Getenv("S3BucketName")
	s3Region := os.Getenv("S3Region")
	s3APIKey := os.Getenv("S3APIKey")
	s3SecretKey := os.Getenv("S3SecretKey")
	s3Domain := os.Getenv("S3Domain")

	s3Provider := uploadprovider.NewS3Provider(s3BucketName, s3Region, s3APIKey, s3SecretKey, s3Domain)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	//run service
	if err := runService(db, s3Provider); err != nil {
		log.Fatalln(err)
	}

}

func runService(db *gorm.DB, upProvider uploadprovider.UploadProvider) error {

	appCtx := component.NewAppContext(db, upProvider)
	r := gin.Default()

	//middleware
	r.Use(middleware.Recover(appCtx))

	// CRUD

	//users
	// users := r.Group("/users")
	// {
	// 	users.POST("", ginuser.CreateUser(appCtx))
	// 	users.GET("/:id", ginuser.GetUser(appCtx))
	// 	users.GET("", ginuser.ListUser(appCtx))
	// 	users.PATCH("/:id", ginuser.UpdateUser(appCtx))
	// 	users.DELETE("/:id", ginuser.DeleteUser(appCtx))
	// }

	//restaurants
	restaurants := r.Group("/restaurants")
	{
		restaurants.POST("", ginrestaurant.CreateRestaurant(appCtx))
		restaurants.GET("/:id", ginrestaurant.GetRestaurant(appCtx))
		restaurants.GET("", ginrestaurant.ListRestaurant(appCtx))
		restaurants.PATCH("/:id", ginrestaurant.UpdateRestaurant(appCtx))
		restaurants.DELETE("/:id", ginrestaurant.DeleteRestaurant(appCtx))
	}
	return r.Run()

}
