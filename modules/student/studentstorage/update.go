package studentstorage

import (
	"context"
	"fmt"
	"studyGoApp/common"
	"studyGoApp/modules/student/studentmodel"
	"time"
)

func (s *sqlStore) UpdateDataByID(ctx context.Context,
	id int,
	data *studentmodel.StudentUpdate,
) error {
	db := s.db
	parsedTime, err := time.Parse("2006-01-02", *data.Birthday)
	if err != nil {
		return err
	}

	fmt.Println(data)

	/* if _, err := db.NamedExec("UPDATE student SET name: name, birthday: birthday WHERE studentID :oldStudentID",
		map[string]interface{}{
			"name":         &data.Name,
			"birthday":     parsedTime,
			"oldStudentID": studentID,
		}); err != nil {
		return err
	} */

	if _, err := db.Exec("UPDATE student SET name = ?, birthday = ? WHERE id = ?", data.Name, parsedTime, id); err != nil {
		return common.ErrDB(err)
	}

	return nil
}
