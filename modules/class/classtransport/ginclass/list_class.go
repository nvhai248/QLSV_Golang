package ginclass

import (
	"net/http"
	"studyGoApp/common"
	"studyGoApp/component"
	"studyGoApp/modules/class/classbiz"
	"studyGoApp/modules/class/classstorage"
	classgrpcclient "studyGoApp/modules/class/classstorage/grpcclient"
	"studyGoApp/proto"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func ListClass(appCtx component.AppContext, cc *grpc.ClientConn) gin.HandlerFunc {
	return func(c *gin.Context) {
		var paging common.Paging

		// Bind query parameters or form data to the filter struct
		if err := c.ShouldBind(&paging); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		paging.Fulfill()

		store := classstorage.NewSQLStore(appCtx.GetMainDBConnection())
		classRegisterStore := classgrpcclient.NewgRPCClient(proto.NewClassRegistrationServiceClient(cc))
		//classRegisterStore := classregisterstorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := classbiz.NewListClassBiz(store, classRegisterStore)
		data, err := biz.ListClass(c.Request.Context(), &paging)

		if err != nil {
			panic(err)
		}

		for i := range data {
			data[i].Mask(false)

			if i == len(data)-1 {
				paging.NextCursor = data[i].FakeID.String()
			}
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(data, paging, nil))
	}
}
