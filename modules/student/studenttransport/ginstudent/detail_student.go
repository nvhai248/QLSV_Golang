package ginstudent

import (
	"net/http"
	"studyGoApp/common"
	"studyGoApp/component"
	"studyGoApp/modules/student/studentbiz"
	"studyGoApp/modules/student/studentstorage"

	"github.com/gin-gonic/gin"
)

func DetailStudent(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		studentID := c.Param("studentID")

		store := studentstorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := studentbiz.NewDetailStudentStore(store)

		result, err := biz.DetailStudent(c.Request.Context(), studentID)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(result))
	}
}
