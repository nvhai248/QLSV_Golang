package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"studyGoApp/component"
	"studyGoApp/component/uploadprovider"
	"studyGoApp/middleware"
	"studyGoApp/modules/class/classtransport/ginclass"
	"studyGoApp/modules/classregister/transport/ginclassregister"
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

func runServices(db *sqlx.DB, secretKey string, upProvider uploadprovider.UploadProvider) {

	appCtx := component.NewAppContext(db, secretKey, upProvider)
	router := gin.Default()

	router.Use(middleware.Recover(appCtx))

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "ping 5555",
		})
	})

	v1 := router.Group("/v1")

	// upload photo
	v1.POST("/upload", ginupload.Upload(appCtx))

	// CRUD student
	// authentication and authorization with JWT
	v1.POST("/students/register", ginstudent.CreateStudent(appCtx))
	v1.POST("/students/login", ginstudent.Login(appCtx))
	students := v1.Group("/students", middleware.RequireAuth(appCtx))
	{
		students.GET("/profile", ginstudent.GetProfile(appCtx))

		students.GET("", ginstudent.ListStudent(appCtx))
		students.GET("/:id", ginstudent.DetailStudent(appCtx))
		students.PATCH("/:id", ginstudent.UpdateStudent(appCtx))
		students.DELETE("/:id", ginstudent.SoftDeleteStudent(appCtx))
	}

	classes := v1.Group("/classes", middleware.RequireAuth(appCtx))
	{
		classes.GET("", ginclass.ListClass(appCtx))
		classes.GET("/:id", ginclass.FindClass(appCtx))
		classes.POST("", ginclass.CreateClass(appCtx))
		classes.DELETE("/:id/cancel_registration", ginclassregister.StudentCancelRegisterClass(appCtx))
		classes.POST("/:id/register", ginclassregister.StudentRegisterClass(appCtx))
		classes.GET("/:id/registered_student", ginclass.GetListRegisteredStudents(appCtx))
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
	secretKey := os.Getenv("SYSTEM_SECRET")
	runServices(db, secretKey, s3upProvider)

}
