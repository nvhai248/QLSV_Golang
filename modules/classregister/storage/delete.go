package classregisterstorage

import (
	"context"
	"studyGoApp/common"
	classregistermodel "studyGoApp/modules/classregister/model"
)

func (s sqlStore) DeleteClassRegister(ctx context.Context, data *classregistermodel.Register) error {
	db := s.db

	if _, err := db.Exec("DELETE FROM class_registers WHERE student_id = ? AND class_id = ?", data.StudentId, data.ClassId); err != nil {
		return common.ErrDB(err)
	}

	return nil
}
