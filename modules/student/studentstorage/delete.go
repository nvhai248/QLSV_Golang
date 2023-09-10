package studentstorage

import (
	"context"
	"studyGoApp/common"
)

func (s *sqlStore) SoftDeleteStudentByStudentID(
	ctx context.Context,
	id int,
) error {
	db := s.db

	if _, err := db.Exec("UPDATE student SET status = ? WHERE id = ?", 0, id); err != nil {
		return common.ErrDB(err)
	}

	return nil
}
