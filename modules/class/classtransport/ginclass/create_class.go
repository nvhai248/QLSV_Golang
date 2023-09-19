package ginclass

import (
	"net/http"
	"studyGoApp/common"
	"studyGoApp/component"
	"studyGoApp/modules/class/classbiz"
	"studyGoApp/modules/class/classmodel"
	"studyGoApp/modules/class/classstorage"

	"github.com/gin-gonic/gin"
)

func CreateClass(appCtx component.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var data classmodel.Class

		if err := ctx.ShouldBindJSON(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		StudentSession := ctx.MustGet(common.CurrentStudent).(common.Requester)
		data.LeaderId = StudentSession.GetId()

		store := classstorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := classbiz.NewCreateClassBiz(store)

		if err := biz.Create(ctx.Request.Context(), &data); err != nil {
			panic(err)
		}

		data.Mask(false)
		ctx.JSON(http.StatusOK, common.SimpleSuccessResponse(data.FakeID.String()))
	}
}
