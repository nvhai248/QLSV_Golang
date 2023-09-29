package classregisterstorage

import (
	"context"
	"database/sql"
	"studyGoApp/common"
	classregistermodel "studyGoApp/modules/classregister/model"
)

func (s sqlStore) FindClassRegister(ctx context.Context, studentId, classId int) (*classregistermodel.Register, error) {
	db := s.db

	var result classregistermodel.Register

	if err := db.Get(&result, "SELECT * FROM class_registers WHERE student_id = ? AND class_id = ?", studentId, classId); err != nil {
		if err == sql.ErrNoRows {
			return nil, common.ErrorNoRows
		}

		return nil, common.ErrDB(err)
	}

	return &result, nil
}
