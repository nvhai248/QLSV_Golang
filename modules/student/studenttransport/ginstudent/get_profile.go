package ginstudent

import (
	"net/http"
	"studyGoApp/common"
	"studyGoApp/component"

	"github.com/gin-gonic/gin"
)

func GetProfile(appCtx component.AppContext) func(*gin.Context) {
	return func(ctx *gin.Context) {
		// get data from context
		data := ctx.MustGet(common.CurrentStudent).(common.Requester)
		//data.Mask(true)

		ctx.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}
}
