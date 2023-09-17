package studentbiz

import (
	"context"

	"studyGoApp/common"
	"studyGoApp/modules/student/studentmodel"
)

type RegisterStorage interface {
	FindByStudentID(ctx context.Context,
		studentID string) (*studentmodel.StudentDetail, error)
	Create(ctx context.Context, data *studentmodel.StudentCreate) error
}

type Hasher interface {
	Hash(data string) string
}

type registerBiz struct {
	store  RegisterStorage
	hasher Hasher
}

func NewRegisterBiz(store RegisterStorage, hasher Hasher) *registerBiz {
	return &registerBiz{store: store, hasher: hasher}
}

func (biz *registerBiz) Register(ctx context.Context, data *studentmodel.StudentCreate) error {

	user, err := biz.store.FindByStudentID(ctx, data.StudentID)

	if err != common.ErrorNoRows && err != nil {
		return err
	}

	if user != nil {
		return common.ErrUsernameExisted
	}

	salt := common.GenSalt(50)

	data.Password = biz.hasher.Hash(data.Password + salt)
	data.Salt = salt
	data.Status = 1
	data.Role = "user" // hard code
	data.Status = 1

	if err := biz.store.Create(ctx, data); err != nil {
		return common.ErrCannotCreateEntity(studentmodel.EntityName, err)
	}

	return nil
}
