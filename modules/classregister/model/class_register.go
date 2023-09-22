package classregistermodel

import (
	"studyGoApp/common"
	"time"
)

type Register struct {
	StudentId int        `json:"student_id", sql:"student_id;"`
	ClassId   int        `json:"class_id", sql:"class_id;"`
	CreatedAt *time.Time `json:"created_at", sql:"created_at;"`

	Students *common.SimpleStudent `json:"students"`
}

func (r Register) TableName() string {
	return "class_registers"
}
