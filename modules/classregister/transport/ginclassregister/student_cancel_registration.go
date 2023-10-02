package ginclassregister

import (
	"net/http"
	"studyGoApp/common"
	"studyGoApp/component"
	"studyGoApp/modules/class/classstorage"
	classregisterbiz "studyGoApp/modules/classregister/biz"
	classregisterstorage "studyGoApp/modules/classregister/storage"
	"studyGoApp/modules/student/studentstorage"

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
		decreaseClassCount := studentstorage.NewSQLStore(appCtx.GetMainDBConnection())
		decreaseStudentCount := classstorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := classregisterbiz.NewCancelRegistrationBiz(store, decreaseClassCount, decreaseStudentCount)

		err = biz.CancelRegistration(ctx.Request.Context(), requester.GetId(), int(uid.GetLocalID()))

		if err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusOK, common.SimpleSuccessResponse(true))

	}
}
