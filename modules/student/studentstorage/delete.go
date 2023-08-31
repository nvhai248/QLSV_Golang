package studentstorage

import (
	"context"
	"studyGoApp/common"
)

func (s *sqlStore) SoftDeleteStudentByStudentID(
	ctx context.Context,
	studentID string,
) error {
	db := s.db

	if _, err := db.Exec("UPDATE student SET status = ? WHERE studentID = ?", 0, studentID); err != nil {
		return common.ErrDB(err)
	}

	return nil
}
