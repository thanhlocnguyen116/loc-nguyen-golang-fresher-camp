package main

import (
	"locnguyen/component"
	"locnguyen/component/uploadprovider"
	"locnguyen/middleware"
	"locnguyen/modules/restaurants/transport/ginrestaurant"
	"locnguyen/modules/upload/transport/ginupload"
	"locnguyen/modules/user/transport/ginuser"
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
	secretKey := os.Getenv("SYSTEM_SECRET")

	s3Provider := uploadprovider.NewS3Provider(s3BucketName, s3Region, s3APIKey, s3SecretKey, s3Domain)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	//run service
	if err := runService(db, s3Provider, secretKey); err != nil {
		log.Fatalln(err)
	}

}

func runService(db *gorm.DB, upProvider uploadprovider.UploadProvider, secretKey string) error {

	appCtx := component.NewAppContext(db, upProvider, secretKey)
	r := gin.Default()

	//middleware
	r.Use(middleware.Recover(appCtx))

	// CRUD

	v := r.Group("/v")

	//users
	v.POST("/register", ginuser.Register(appCtx))
	v.POST("/login", ginuser.Login(appCtx))
	v.GET("/profile", middleware.RequiredAuth(appCtx), ginuser.GetProfile(appCtx))

	//upload images
	v.POST("/upload", ginupload.Upload(appCtx))

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
