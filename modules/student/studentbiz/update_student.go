package studentbiz

import (
	"context"
	"studyGoApp/common"
	"studyGoApp/modules/student/studentmodel"
)

type UpdateStudentStore interface {
	DetailStudent(
		ctx context.Context,
		id int,
	) (*studentmodel.StudentDetail, error)
	UpdateDataByID(
		ctx context.Context,
		id int,
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
	id int,
	data *studentmodel.StudentUpdate,
) error {
	oldData, err := biz.store.DetailStudent(ctx, id)

	if err != nil {
		return common.ErrCannotGetEntity(studentmodel.EntityName, err)
	}

	if oldData.Status == 0 {
		return common.NewCustomError(nil, "Data deleted!", studentmodel.EntityName)
	}

	if data.Name == nil {
		data.Name = &oldData.Name
	}

	if data.Birthday == nil {
		data.Birthday = &oldData.Birthday
	}

	if err = biz.store.UpdateDataByID(ctx, id, data); err != nil {
		return common.ErrCannotUpdateEntity(studentmodel.EntityName, err)
	}

	return nil
}
