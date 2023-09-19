package classmodel

import (
	"studyGoApp/common"
)

const EntityName = "class"

type Class struct {
	common.SQLModel `json:", inline"`
	Name            string `json:"name" db:"name"`
	LeaderId        int    `json:"leaderId" db:"leaderId"`
	SchoolYearStart int    `json:"school_year_start" db:"school_year_start"`
	SchoolYearEnd   int    `json:"school_year_end" db:"school_year_end"`
	Status          int    `json:"status" db:"status"`
	Semester        int    `json:"semester" db:"semester"`
}

func (Class) TableName() string {
	return "classes"
}

func (c *Class) Mask(isAdminOrOwner bool) {
	c.GenUID(common.DbTypeClass)
}
