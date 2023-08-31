package studentstorage

import (
	"context"
	"database/sql"
	"studyGoApp/common"
	"studyGoApp/modules/student/studentmodel"
)

func (s *sqlStore) DetailStudent(ctx context.Context,
	studentID string,
) (*studentmodel.StudentDetail, error) {
	db := s.db

	var detailOfStudent studentmodel.StudentDetail

	if err := db.Get(&detailOfStudent, "SELECT * FROM student WHERE studentID = ?", studentID); err != nil {
		if err == sql.ErrNoRows {
			return nil, common.ErrorNoRows
		}
		return nil, common.ErrDB(err)
	}

	return &detailOfStudent, nil
}
