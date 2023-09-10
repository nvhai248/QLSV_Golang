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
		uid, err := common.FromBase58(c.Param("id"))

		if err != nil {
			panic(common.ErrInternal(err))
		}

		store := studentstorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := studentbiz.NewSoftDeleteStudentBiz(store)

		if err := biz.SoftDeleteStudent(c.Request.Context(), int(uid.GetLocalID())); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
