package studentstorage

import (
	"context"
	"studyGoApp/common"
	"studyGoApp/modules/student/studentmodel"
)

func (s *sqlStore) ListDataByCondition(ctx context.Context,
	conditions map[string]interface{},
	filter *studentmodel.Filter,
	paging *common.Paging,
	moreKeys ...string,
) ([]studentmodel.Student, error) {
	db := s.db

	args := []interface{}{}

	query := "SELECT * FROM student"

	var conditionsAndMore string

	// add conditions
	i := 0
	if len(conditions) > 0 {
		conditionsAndMore += " WHERE "

		for key, value := range conditions {
			if i > 0 {
				conditionsAndMore += " AND "
			}
			conditionsAndMore += key + " = ? "
			i++

			args = append(args, value)
		}
	}

	// add filter conditions
	if v := filter.Name; v != "" {
		if len(conditions) > 0 {
			conditionsAndMore += " AND "
		} else {
			conditionsAndMore += " WHERE "
		}
		conditionsAndMore += "name = ?"
		args = append(args, v)
	}

	var students []studentmodel.Student

	offset := (paging.Page - 1) * paging.Limit
	limit := paging.Limit

	conditionsAndMore = conditionsAndMore + " ORDER BY id DESC LIMIT ? OFFSET ?"
	args = append(args, limit, offset)

	query = db.Rebind(query + conditionsAndMore)
	if err := db.Select(&students, query, args...); err != nil {
		return nil, err
	}

	// count paging
	var total int64
	countQuery := "SELECT COUNT(*) FROM student"

	if err := db.Get(&total, countQuery); err != nil {
		return nil, err
	}

	paging.Total = total

	return students, nil
}
