package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"studyGoApp/component"
	"studyGoApp/component/uploadprovider"
	"studyGoApp/middleware"
	"studyGoApp/modules/student/studenttransport/ginstudent"
	"studyGoApp/modules/upload/uploadtransport/ginupload"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
)

func ConnectToDB(dns string) *sqlx.DB {
	db, err := sqlx.Connect("mysql", dns)
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}
	// Check connection to DB by Ping
	err = db.Ping()
	if err != nil {
		log.Fatal("Failed to ping the database:", err)
	}
	fmt.Println("Connected to the database!")

	return db
}

func runServices(db *sqlx.DB, upProvider uploadprovider.UploadProvider) {

	appCtx := component.NewAppContext(db, upProvider)
	router := gin.Default()

	router.Use(middleware.Recover(appCtx))

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "ping 5555",
		})
	})

	// upload photo
	router.POST("/upload", ginupload.Upload(appCtx))

	// CRUD
	students := router.Group("/students")
	{
		students.GET("", ginstudent.ListStudent(appCtx))
		students.GET("/:studentID", ginstudent.DetailStudent(appCtx))
		students.POST("/", ginstudent.CreateStudent(appCtx))
		students.PATCH("/:studentID", ginstudent.UpdateStudent(appCtx))
		students.DELETE("/:studentID", ginstudent.SoftDeleteStudent(appCtx))
	}

	router.Run(":8080")
}

func main() {
	// Load the .env file
	if err := godotenv.Load(); err != nil {
		log.Fatalln(err)
	}

	dns := os.Getenv("DBConnectionStr")

	// Connect to the database
	db := ConnectToDB(dns)

	s3BucketName := os.Getenv("S3BucketName")
	s3Region := os.Getenv("S3Region")
	s3ApiKey := os.Getenv("S3ApiKey")
	s3Secret := os.Getenv("S3Secret")
	s3Domain := os.Getenv("S3Domain")
	s3upProvider := uploadprovider.NewS3Provider(s3BucketName, s3Region, s3ApiKey, s3Secret, s3Domain)

	runServices(db, s3upProvider)
}
