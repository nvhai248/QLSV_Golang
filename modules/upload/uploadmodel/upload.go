package uploadmodel

import "studyGoApp/common"

type ImageUpload struct {
	common.Image `json:", inline"`
}

func (ImageUpload) TableName() string { return "images" }

func ErrCannotSaveFile(err error) *common.AppError {
	return common.NewErrorResponse(err, "Cannot Save file!", err.Error(), "FILE_ERROR")
}

func ErrFileIsNotImage(err error) *common.AppError {
	return common.NewErrorResponse(err, "File is not Image!", err.Error(), "FILE_ERROR")
}
