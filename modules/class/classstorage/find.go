package classstorage

import (
	"context"
	"studyGoApp/common"
	"studyGoApp/modules/class/classmodel"
)

func (s *sqlStore) FindClassById(ctx context.Context, id int, condition map[string]interface{}) (*classmodel.Class, error) {
	db := s.db

	var result classmodel.Class

	query := "SELECT * FROM classes WHERE id = ? "
	args := []interface{}{}

	args = append(args, id)
	// add condition to the query
	if len(condition) > 0 {
		for key, value := range condition {
			query += "AND" + key + " = ? "
			args = append(args, value)
		}
	}

	query = db.Rebind(query)
	if err := db.Get(&result, query, args...); err != nil {
		return nil, common.ErrDB(err)
	}

	return &result, nil
}
