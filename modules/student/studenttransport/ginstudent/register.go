package ginstudent

import (
	"net/http"
	"studyGoApp/common"
	"studyGoApp/component"
	"studyGoApp/component/hasher"
	"studyGoApp/modules/student/studentbiz"
	"studyGoApp/modules/student/studentmodel"
	"studyGoApp/modules/student/studentstorage"

	"github.com/gin-gonic/gin"
)

func CreateStudent(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data studentmodel.StudentCreate

		if err := c.ShouldBindJSON(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := studentstorage.NewSQLStore(appCtx.GetMainDBConnection())
		md5 := hasher.NewMd5Hash()
		biz := studentbiz.NewRegisterBiz(store, md5)

		if err := biz.Register(c.Request.Context(), &data); err != nil {
			panic(err)
		}

		data.Mask(false)

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data.FakeID.String()))
	}
}
