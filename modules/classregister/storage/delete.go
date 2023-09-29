package classregisterstorage

import (
	"context"
	"studyGoApp/common"
)

func (s sqlStore) DeleteClassRegister(ctx context.Context, studentId, classId int) error {
	db := s.db

	if _, err := db.Exec("DELETE FROM class_registers WHERE student_id = ? AND class_id = ?", studentId, classId); err != nil {
		return common.ErrDB(err)
	}

	return nil
}
