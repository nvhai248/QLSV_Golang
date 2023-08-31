package studentbiz

import (
	"context"

	"studyGoApp/common"
	"studyGoApp/modules/student/studentmodel"
)

type CreateStudentStore interface {
	Create(ctx context.Context, data *studentmodel.StudentCreate) error
}

type createStudentBiz struct {
	store CreateStudentStore
}

func NewCreateStudentBiz(store CreateStudentStore) *createStudentBiz {
	return &createStudentBiz{store: store}
}

func (biz *createStudentBiz) CreateStudent(ctx context.Context, data *studentmodel.StudentCreate) error {

	if err := data.Validate(); err != nil {
		return common.ErrCannotCreateEntity(studentmodel.EntityName, err)
	}

	err := biz.store.Create(ctx, data)

	return err
}
