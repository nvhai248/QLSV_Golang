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
		uid, err := common.FromBase58(c.Param("id"))

		if err != nil {
			panic(common.ErrInternal(err))
		}

		store := studentstorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := studentbiz.NewDetailStudentStore(store)

		result, err := biz.DetailStudent(c.Request.Context(), int(uid.GetLocalID()))

		if err != nil {
			panic(err)
		}

		result.Mask(false)

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(result))
	}
}
