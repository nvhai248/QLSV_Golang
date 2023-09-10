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
		uid, err := common.FromBase58(c.Param("id"))

		if err != nil {
			panic(common.ErrInternal(err))
		}

		var data studentmodel.StudentUpdate

		if err := c.ShouldBindJSON(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := studentstorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := studentbiz.NewUpdateStudentBiz(store)

		if err := biz.UpdateStudent(c.Request.Context(), int(uid.GetLocalID()), &data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
