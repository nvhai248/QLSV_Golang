package studentbiz

import (
	"context"
	"errors"
	"studyGoApp/modules/student/studentmodel"
)

type SoftDeleteStudentStore interface {
	DetailStudent(
		ctx context.Context,
		studentID string,
	) (*studentmodel.StudentDetail, error)
	SoftDeteleStudentByStudentID(
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
		return err
	}

	if oldData.Status == 0 {
		return errors.New("Data deleted!")
	}

	if err = biz.store.SoftDeteleStudentByStudentID(ctx, oldData.StudentID); err != nil {
		return err
	}

	return nil
}
