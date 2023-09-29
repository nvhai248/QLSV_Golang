package ginclassregister

import (
	"net/http"
	"studyGoApp/common"
	"studyGoApp/component"
	classregisterbiz "studyGoApp/modules/classregister/biz"
	classregistermodel "studyGoApp/modules/classregister/model"
	classregisterstorage "studyGoApp/modules/classregister/storage"

	"github.com/gin-gonic/gin"
)

//POST /v1/classes/:id/register_class

func StudentRegisterClass(appCtx component.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		uid, err := common.FromBase58(ctx.Param("id"))

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		requester := ctx.MustGet(common.CurrentStudent).(common.Requester)
		data := &classregistermodel.Register{
			StudentId: requester.GetId(),
			ClassId:   int(uid.GetLocalID()),
		}

		store := classregisterstorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := classregisterbiz.NewRegisterBiz(store)

		err = biz.Register(ctx.Request.Context(), data)

		if err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusOK, common.SimpleSuccessResponse(data))

	}
}
