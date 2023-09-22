package classregisterstorage

import (
	"context"
	"fmt"
	"studyGoApp/common"
	classregistermodel "studyGoApp/modules/classregister/model"
)

func (s sqlStore) GetStudentRegister(ctx context.Context, ids []int) (map[int]int, error) {
	result := make(map[int]int)
	db := s.db
	for _, id := range ids {
		count := 0
		if err := db.Get(&count, "SELECT COUNT(*) FROM class_registers WHERE student_id = ?", id); err != nil {
			return nil, common.ErrDB(err)
		}

		result[id] = count
	}

	return result, nil
}

func (s sqlStore) GetNumberOfStudentRegisteredInClass(ctx context.Context, ids []int) (map[int]int, error) {
	result := make(map[int]int)
	db := s.db
	for _, id := range ids {
		count := 0
		if err := db.Get(&count, "SELECT COUNT(*) FROM class_registers WHERE class_id = ?", id); err != nil {
			return nil, common.ErrDB(err)
		}

		result[id] = count
	}

	return result, nil
}

func (s sqlStore) GetListSimpleStudentByConditions(ctx context.Context,
	conditions map[string]interface{},
	filter *classregistermodel.Filter,
	paging *common.Paging,
) ([]common.SimpleStudent, error) {
	db := s.db

	args := []interface{}{}

	query := "SELECT student.id, student.studentID, student.name, student.birthday, student.role, student.created_at, student.updated_at FROM student"

	// add filter conditions
	if v := filter.ClassId; v != 0 {
		query += " JOIN class_registers ON class_registers.student_id = student.id WHERE class_registers.class_id = ?"
		args = append(args, v)
	}

	// add conditions
	if len(conditions) > 0 {
		for key, value := range conditions {
			query += " AND " + key + " = ? "
			args = append(args, value)
		}
	}

	var students []common.SimpleStudent
	limit := paging.Limit

	// updated paging
	if v := paging.FakeCursor; v != "" {
		query = query + fmt.Sprintf(" AND class_registers.created_at < '%s' ", paging.FakeCursor) + "ORDER BY id DESC LIMIT ?"
		args = append(args, limit)
	} else {
		offset := (paging.Page - 1) * paging.Limit

		query = query + " ORDER BY student.id DESC LIMIT ? OFFSET ?"
		args = append(args, limit, offset)
	}

	fmt.Println(query)
	fmt.Println(args...)

	query = db.Rebind(query)
	if err := db.Select(&students, query, args...); err != nil {
		return nil, common.ErrDB(err)
	}

	// count paging
	paging.Total = int64(len(students))

	return students, nil
}
