package studentstorage

import (
	"context"
	studentmodel "studyGoApp/modules/student/studentmodel"
	"time"
)

func (s *sqlStore) Create(ctx context.Context, data *studentmodel.StudentCreate) error {
	db := s.db

	parsedTime, err := time.Parse("2006-01-02", data.Birthday)
	if err != nil {
		return err
	}

	if _, err := db.Exec("INSERT INTO student (name, studentID, birthday, status) VALUES (?, ?, ?, ?)", data.Name, data.StudentID, parsedTime, 1); err != nil {
		return err
	}

	return nil
}
