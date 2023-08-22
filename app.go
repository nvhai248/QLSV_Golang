package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"studyGoApp/component"
	"studyGoApp/modules/student/studenttransport/ginstudent"

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

func APIs(db *sqlx.DB) {
	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "ping 5555",
		})
	})

	// CRUD
	appCtx := component.NewAppContext(db)

	students := router.Group("/students")
	{
		students.GET("", ginstudent.ListStudent(appCtx))
		students.GET("/:studentID", ginstudent.DetailStudent(appCtx))
		students.POST("/add", ginstudent.CreateStudent(appCtx))

		// Other routes...

		students.PATCH("/:studentID", ginstudent.UpdateStudent(appCtx))

		students.DELETE("/:id", func(ctx *gin.Context) {
			id := ctx.Param("id")

			if _, err := db.Exec("DELETE FROM student WHERE studentID = ?", id); err != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}

			ctx.JSON(http.StatusOK, gin.H{"message": "Success delete!"})
		})
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

	fmt.Println(db)

	APIs(db)
}
