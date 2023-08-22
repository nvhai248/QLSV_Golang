package studentbiz

import (
	"context"
	"studyGoApp/modules/student/studentmodel"
)

type DetailStudentStore interface {
	DetailStudent(ctx context.Context,
		studentID string,
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
	studentID string,
) (*studentmodel.StudentDetail, error) {
	result, err := s.store.DetailStudent(ctx, studentID)

	return result, err
}
