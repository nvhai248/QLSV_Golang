package classstorage

import (
	"context"
	"studyGoApp/common"
)

func (s *sqlStore) IncreaseStudentCount(ctx context.Context, id int) error {
	db := s.db

	if _, err := db.Exec("UPDATE classes SET student_count = student_count + 1 WHERE id = ?", id); err != nil {
		return common.ErrDB(err)
	}

	return nil
}

func (s *sqlStore) DecreaseStudentCount(ctx context.Context, id int) error {
	db := s.db

	if _, err := db.Exec("UPDATE classes SET student_count = student_count - 1 WHERE id = ?", id); err != nil {
		return common.ErrDB(err)
	}

	return nil
}
