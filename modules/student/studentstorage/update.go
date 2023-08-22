package studentstorage

import (
	"context"
	"fmt"
	"studyGoApp/modules/student/studentmodel"
	"time"
)

func (s *sqlStore) UpdateDataByID(ctx context.Context,
	studentID string,
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

	if _, err := db.Exec("UPDATE student SET name = ?, birthday = ? WHERE studentID = ?", data.Name, parsedTime, studentID); err != nil {
		return err
	}

	return nil
}
