package ginstudent

import (
	"net/http"
	"studyGoApp/common"
	"studyGoApp/component"
	"studyGoApp/modules/student/studentbiz"
	"studyGoApp/modules/student/studentstorage"

	"github.com/gin-gonic/gin"
)

func SoftDeleteStudent(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		studentID := c.Param("studentID")

		store := studentstorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := studentbiz.NewSoftDeleteStudentBiz(store)

		if err := biz.SoftDeleteStudent(c.Request.Context(), studentID); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
