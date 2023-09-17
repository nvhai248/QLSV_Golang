package studentstorage

import (
	"context"
	"studyGoApp/common"
	studentmodel "studyGoApp/modules/student/studentmodel"
	"time"
)

func (s *sqlStore) Create(ctx context.Context, data *studentmodel.StudentCreate) error {
	db := s.db

	parsedTime, err := time.Parse("2006-01-02", data.Birthday)
	if err != nil {
		return common.ErrInvalidRequest(err)
	}

	if _, err := db.Exec(
		"INSERT INTO student (name, studentID, birthday, status, avatar, cover, password, fb_id, gg_id, salt, role) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)",
		data.Name, data.StudentID, parsedTime, 1, data.Avatar, data.Cover, data.Password, data.FbId, data.GgId, data.Salt, data.Role); err != nil {
		return common.ErrDB(err)
	}

	return nil
}
