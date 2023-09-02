package uploadbiz

import (
	"context"
	"fmt"
	"path/filepath"
	"strings"
	"studyGoApp/common"
	"studyGoApp/component/uploadprovider"
	"studyGoApp/modules/upload/uploadmodel"
	"time"
)

type CreateImageStorage interface {
	CreateImage(context context.Context, data *common.Image) error
}

type uploadBiz struct {
	provider uploadprovider.UploadProvider
	imgStore CreateImageStorage
}

func NewUploadBiz(provider uploadprovider.UploadProvider, imgStore CreateImageStorage) *uploadBiz {
	return &uploadBiz{provider: provider, imgStore: imgStore}
}

func (biz *uploadBiz) Upload(ctx context.Context, data []byte, folder, fileName string) (*common.Image, error) {
	/* reader := bytes.NewReader(data)

	IMG, _, err := image.DecodeConfig(reader)

	if err != nil {
		return nil, uploadmodel.ErrFileIsNotImage(err)
	} */

	w := 0
	h := 0

	if strings.TrimSpace(folder) == "" {
		folder = "img"
	}

	fileExt := filepath.Ext(fileName)                                // "img.jpg" -> "jpg"
	fileName = fmt.Sprintf("%d%s", time.Now().Nanosecond(), fileExt) // 9128213314.jpg

	img, err := biz.provider.SaveFileUploaded(ctx, data, fmt.Sprintf("%s/%s", folder, fileName))

	if err != nil {
		return nil, uploadmodel.ErrCannotSaveFile(err)
	}

	img.Width = w
	img.Height = h
	img.CloudName = "s3" //should be set in provider
	img.Extension = fileExt

	if err := biz.imgStore.CreateImage(ctx, img); err != nil {
		// delete image S3
		return nil, uploadmodel.ErrCannotSaveFile(err)
	}

	return img, nil
}
