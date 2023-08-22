package studentbiz

import (
	"context"
	"errors"
	"studyGoApp/modules/student/studentmodel"
)

type UpdateStudentStore interface {
	DetailStudent(
		ctx context.Context,
		studentID string,
	) (*studentmodel.StudentDetail, error)
	UpdateDataByID(
		ctx context.Context,
		studentID string,
		data *studentmodel.StudentUpdate,
	) error
}

type updateStudentBiz struct {
	store UpdateStudentStore
}

func NewUpdateStudentBiz(store UpdateStudentStore) *updateStudentBiz {
	return &updateStudentBiz{
		store: store,
	}
}

func (biz *updateStudentBiz) UpdateStudent(ctx context.Context,
	studentID string,
	data *studentmodel.StudentUpdate,
) error {
	oldData, err := biz.store.DetailStudent(ctx, studentID)

	if err != nil {
		return err
	}

	if oldData.Status == 0 {
		return errors.New("Data deleted!")
	}

	if err = biz.store.UpdateDataByID(ctx, oldData.StudentID, data); err != nil {
		return err
	}

	return nil
}
