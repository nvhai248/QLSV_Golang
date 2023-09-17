package studentstorage

import (
	"context"
	"database/sql"
	"studyGoApp/common"
	"studyGoApp/modules/student/studentmodel"
)

func (s *sqlStore) FindByStudentID(ctx context.Context,
	studentID string) (*studentmodel.StudentDetail, error) {
	db := s.db

	var result studentmodel.StudentDetail

	if err := db.Get(&result, "SELECT * FROM student WHERE studentID = ?", studentID); err != nil {
		if err == sql.ErrNoRows {
			return nil, common.ErrorNoRows
		}
		return nil, common.ErrDB(err)
	}

	return &result, nil
}
