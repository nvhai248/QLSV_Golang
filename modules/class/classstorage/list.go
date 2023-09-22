package classstorage

import (
	"context"
	"fmt"
	"studyGoApp/common"
	"studyGoApp/modules/class/classmodel"
)

func (s *sqlStore) ListClassByCondition(ctx context.Context,
	conditions map[string]interface{},
	paging *common.Paging,
	moreKeys ...string,
) ([]classmodel.Class, error) {
	db := s.db

	args := []interface{}{}

	query := "SELECT * FROM classes"

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

	var classes []classmodel.Class
	limit := paging.Limit

	// updated paging
	if v := paging.FakeCursor; v != "" {
		if i == 0 {
			conditionsAndMore += " WHERE "
		} else {
			conditionsAndMore += " AND "
		}

		if uid, err := common.FromBase58(v); err == nil {
			conditionsAndMore = conditionsAndMore + fmt.Sprintf("id < %d ", int(uid.GetLocalID())) + "ORDER BY id DESC LIMIT ?"
			args = append(args, limit)
		}
	} else {
		offset := (paging.Page - 1) * paging.Limit

		conditionsAndMore = conditionsAndMore + " ORDER BY classes.id DESC LIMIT ? OFFSET ?"
		args = append(args, limit, offset)
	}

	query = db.Rebind(query + conditionsAndMore)
	if err := db.Select(&classes, query, args...); err != nil {
		return nil, common.ErrDB(err)
	}

	for i := 0; i < len(classes); i++ {
		var simpleStudent common.SimpleStudent
		if err := db.Get(&simpleStudent, "SELECT id, studentID, name, studentID, birthday, role, created_at, updated_at FROM student WHERE id = ?",
			classes[i].LeaderId); err != nil {
			return nil, common.ErrDB(err)
		}
		simpleStudent.Mask(false)
		classes[i].ClassMonitor = &simpleStudent
	}

	// count paging
	var total int64
	countQuery := "SELECT COUNT(*) FROM classes WHERE status in (1)"

	if err := db.Get(&total, countQuery); err != nil {
		return nil, common.ErrDB(err)
	}

	paging.Total = total

	return classes, nil
}
