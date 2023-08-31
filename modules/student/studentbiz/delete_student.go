package studentbiz

import (
	"context"
	"studyGoApp/common"
	"studyGoApp/modules/student/studentmodel"
)

type SoftDeleteStudentStore interface {
	DetailStudent(
		ctx context.Context,
		studentID string,
	) (*studentmodel.StudentDetail, error)
	SoftDeleteStudentByStudentID(
		ctx context.Context,
		studentID string,
	) error
}

type softDeleteStudentBiz struct {
	store SoftDeleteStudentStore
}

func NewSoftDeleteStudentBiz(store SoftDeleteStudentStore) *softDeleteStudentBiz {
	return &softDeleteStudentBiz{
		store: store,
	}
}

func (biz *softDeleteStudentBiz) SoftDeleteStudent(ctx context.Context,
	studentID string,
) error {
	oldData, err := biz.store.DetailStudent(ctx, studentID)

	if err != nil {
		return common.ErrCannotGetEntity(studentmodel.EntityName, err)
	}

	if oldData.Status == 0 {
		return common.NewCustomError(nil, "Data deleted!", studentmodel.EntityName)
	}

	if err = biz.store.SoftDeleteStudentByStudentID(ctx, oldData.StudentID); err != nil {
		return common.ErrCannotDeleteEntity(studentmodel.EntityName, err)
	}

	return nil
}
