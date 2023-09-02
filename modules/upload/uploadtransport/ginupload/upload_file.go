package ginupload

import (
	"net/http"
	"studyGoApp/common"
	"studyGoApp/component"
	"studyGoApp/modules/upload/uploadbiz"
	"studyGoApp/modules/upload/uploadstorage"

	"github.com/gin-gonic/gin"
)

func Upload(appCtx component.AppContext) func(*gin.Context) {
	return func(c *gin.Context) {
		db := appCtx.GetMainDBConnection()

		fileHeader, err := c.FormFile("file")

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		folder := c.DefaultPostForm("folder", "img")

		file, err := fileHeader.Open()

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		defer file.Close() // we can close here

		dataBytes := make([]byte, fileHeader.Size)
		if _, err := file.Read(dataBytes); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		imgStore := uploadstorage.NewSQLStore(db)
		biz := uploadbiz.NewUploadBiz(appCtx.UploadProvider(), imgStore)

		img, err := biz.Upload(c.Request.Context(), dataBytes, folder, fileHeader.Filename)

		if err != nil {
			panic(err)
		}

		//upload in root directory (server)
		//c.SaveUploadedFile(fileHeader, fmt.Sprintf("./static/%s", fileHeader.Filename))

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(img))
	}
}
