package classstorage

import (
	"context"
	"studyGoApp/common"
	"studyGoApp/modules/class/classmodel"
)

func (s *sqlStore) CreateClass(ctx context.Context, data *classmodel.Class) error {
	db := s.db

	if _, err := db.Exec("INSERT INTO classes (name, leaderId, school_year_start, school_year_end, status, semester) VALUES (?, ?, ?, ?, ?, ?)",
		data.Name, data.LeaderId, data.SchoolYearStart, data.SchoolYearEnd, data.Status, data.Semester); err != nil {
		return common.ErrDB(err)
	}

	return nil
}
