package ginstudent

import (
	"net/http"
	"studyGoApp/common"
	"studyGoApp/component"
	classregisterstorage "studyGoApp/modules/classregister/storage"
	"studyGoApp/modules/student/studentbiz"
	"studyGoApp/modules/student/studentmodel"
	"studyGoApp/modules/student/studentstorage"

	"github.com/gin-gonic/gin"
)

func ListStudent(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var filter studentmodel.Filter

		// Bind query parameters or form data to the filter struct
		if err := c.ShouldBind(&filter); err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		var paging common.Paging

		// Bind query parameters or form data to the filter struct
		if err := c.ShouldBind(&paging); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		paging.Fulfill()
		/* fmt.Println(paging.Page)
		fmt.Println(paging.Limit) */

		store := studentstorage.NewSQLStore(appCtx.GetMainDBConnection())
		classStore := classregisterstorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := studentbiz.NewListStudentBiz(store, classStore)

		data, err := biz.ListStudent(c.Request.Context(), &filter, &paging)

		if err != nil {
			panic(err)
		}

		for i := range data {
			data[i].Mask(false)

			if i == len(data)-1 {
				paging.NextCursor = data[i].FakeID.String()
			}
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(data, paging, filter))
	}
}
