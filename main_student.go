package main

import (
	"studyGoApp/component"
	"studyGoApp/middleware"
	"studyGoApp/modules/student/studenttransport/ginstudent"

	"github.com/gin-gonic/gin"
)

func MainStudent(appCtx component.AppContext, ver *gin.RouterGroup) {
	students := ver.Group("/students", middleware.RequireAuth(appCtx))
	{
		students.GET("/profile", ginstudent.GetProfile(appCtx))

		students.GET("", ginstudent.ListStudent(appCtx))
		students.GET("/:id", ginstudent.DetailStudent(appCtx))
		students.PATCH("/:id", ginstudent.UpdateStudent(appCtx))
		students.DELETE("/:id", ginstudent.SoftDeleteStudent(appCtx))
	}
}
