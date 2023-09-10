package studentbiz

import (
	"context"
	"studyGoApp/common"
	"studyGoApp/modules/student/studentmodel"
)

type DetailStudentStore interface {
	DetailStudent(ctx context.Context,
		id int,
	) (*studentmodel.StudentDetail, error)
}

type detailStudentStore struct {
	store DetailStudentStore
}

func NewDetailStudentStore(store DetailStudentStore) *detailStudentStore {
	return &detailStudentStore{
		store: store,
	}
}

func (s *detailStudentStore) DetailStudent(ctx context.Context,
	id int,
) (*studentmodel.StudentDetail, error) {
	result, err := s.store.DetailStudent(ctx, id)

	if err != nil {
		if err != common.ErrorNoRows {
			return nil, common.ErrCannotGetEntity(studentmodel.EntityName, err)
		}

		return nil, common.ErrCannotGetEntity(studentmodel.EntityName, err)
	}

	if result.Status != 1 {
		return nil, common.ErrEntityDeleted(studentmodel.EntityName, nil)
	}

	return result, err
}
