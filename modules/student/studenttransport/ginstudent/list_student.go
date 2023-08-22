package ginstudent

import (
	"net/http"
	"studyGoApp/common"
	"studyGoApp/component"
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
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		var paging common.Paging

		// Bind query parameters or form data to the filter struct
		if err := c.ShouldBind(&paging); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		paging.Fulfill()
		/* fmt.Println(paging.Page)
		fmt.Println(paging.Limit) */

		store := studentstorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := studentbiz.NewListStudentBiz(store)

		data, err := biz.ListStudent(c.Request.Context(), &filter, &paging)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(data, paging, filter))
	}
}
