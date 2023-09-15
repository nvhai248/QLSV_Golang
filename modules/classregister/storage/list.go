package classregisterstorage

import (
	"context"
	"studyGoApp/common"
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
