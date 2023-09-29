package ginclassregister

import (
	"net/http"
	"studyGoApp/common"
	"studyGoApp/component"
	classregisterbiz "studyGoApp/modules/classregister/biz"
	classregisterstorage "studyGoApp/modules/classregister/storage"

	"github.com/gin-gonic/gin"
)

//POST /v1/classes/:id/register_class

func StudentCancelRegisterClass(appCtx component.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		uid, err := common.FromBase58(ctx.Param("id"))

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		requester := ctx.MustGet(common.CurrentStudent).(common.Requester)

		store := classregisterstorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := classregisterbiz.NewCancelRegistrationBiz(store)

		err = biz.CancelRegistration(ctx.Request.Context(), requester.GetId(), int(uid.GetLocalID()))

		if err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusOK, common.SimpleSuccessResponse(true))

	}
}
