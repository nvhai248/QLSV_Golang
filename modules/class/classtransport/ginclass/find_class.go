package ginclass

import (
	"net/http"
	"studyGoApp/common"
	"studyGoApp/component"
	"studyGoApp/modules/class/classbiz"
	"studyGoApp/modules/class/classstorage"

	"github.com/gin-gonic/gin"
)

func FindClass(appCtx component.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		uid, err := common.FromBase58(ctx.Param("id"))

		if err != nil {
			panic(common.ErrInternal(err))
		}

		store := classstorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := classbiz.NewFindClassBiz(store)

		result, err := biz.FindClassById(ctx.Request.Context(), int(uid.GetLocalID()), nil)

		if err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusOK, common.SimpleSuccessResponse(result))
	}
}
