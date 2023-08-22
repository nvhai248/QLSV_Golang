package ginstudent

import (
	"net/http"
	"studyGoApp/common"
	"studyGoApp/component"
	"studyGoApp/modules/student/studentbiz"
	"studyGoApp/modules/student/studentmodel"
	"studyGoApp/modules/student/studentstorage"

	"github.com/gin-gonic/gin"
)

func UpdateStudent(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		studentID := c.Param("studentID")

		var data studentmodel.StudentUpdate

		if err := c.ShouldBindJSON(&data); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		store := studentstorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := studentbiz.NewUpdateStudentBiz(store)

		if err := biz.UpdateStudent(c.Request.Context(), studentID, &data); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(data, nil, nil))
	}
}
