package ginclass

import (
	"net/http"
	"studyGoApp/common"
	"studyGoApp/component"
	"studyGoApp/modules/class/classbiz"
	classregistermodel "studyGoApp/modules/classregister/model"
	classregisterstorage "studyGoApp/modules/classregister/storage"

	"github.com/gin-gonic/gin"
)

func GetListRegisteredStudents(appCtx component.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		uid, err := common.FromBase58(ctx.Param("id"))
		if err != nil {
			panic(common.ErrInternal(err))
		}
		var paging common.Paging

		// Bind query parameters or form data to the filter struct
		if err := ctx.ShouldBind(&paging); err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		paging.Fulfill()

		var filter classregistermodel.Filter

		filter.ClassId = int(uid.GetLocalID())

		store := classregisterstorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := classbiz.NewListRegisteredStuBiz(store)

		result, err := biz.GetListRegisteredStu(ctx.Request.Context(), &filter, &paging)

		if err != nil {
			panic(err)
		}

		for i := range result {
			result[i].Mask(false)

			if i == len(result)-1 {
				paging.NextCursor = result[i].CreatedAt
			}
		}

		filter.Mask()

		ctx.JSON(http.StatusOK, common.NewSuccessResponse(result, paging, filter))
	}
}
