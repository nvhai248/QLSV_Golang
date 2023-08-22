package studentstorage

import (
	"context"
	"studyGoApp/modules/student/studentmodel"
)

func (s *sqlStore) DetailStudent(ctx context.Context,
	studentID string,
) (*studentmodel.StudentDetail, error) {
	db := s.db
	var detailOfStudent studentmodel.StudentDetail
	if err := db.Get(&detailOfStudent, "SELECT * FROM student WHERE studentID = ?", studentID); err != nil {
		return nil, err
	}

	return &detailOfStudent, nil
}
