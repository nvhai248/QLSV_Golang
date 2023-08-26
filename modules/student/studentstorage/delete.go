package studentstorage

import (
	"context"
)

func (s *sqlStore) SoftDeteleStudentByStudentID(
	ctx context.Context,
	studentID string,
) error {
	db := s.db

	if _, err := db.Exec("UPDATE student SET status = ? WHERE studentID = ?", 0, studentID); err != nil {
		return err
	}

	return nil
}
