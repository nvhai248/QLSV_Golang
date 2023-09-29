package classregisterstorage

import (
	"context"
	"fmt"
	"studyGoApp/common"
	classregistermodel "studyGoApp/modules/classregister/model"
)

func (s sqlStore) CreateClassRegister(ctx context.Context, data *classregistermodel.Register) error {
	db := s.db

	fmt.Print("Ad", data.StudentId, data.ClassId)

	if _, err := db.Exec("INSERT INTO class_registers (student_id, class_id) VALUES (?, ?)", data.StudentId, data.ClassId); err != nil {
		return common.ErrDB(err)
	}

	return nil
}
