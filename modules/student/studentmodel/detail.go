package studentmodel

import (
	"studyGoApp/common"
)

type StudentDetail struct {
	common.SQLModel `json: ", inline"`
	StudentID       string         `json:"studentID" db:"studentID"`
	Name            string         `json:"name" db:"name"`
	Birthday        string         `json:"birthday" db:"birthday"`
	Avatar          *common.Image  `json:"avatar" json:"avatar"`
	Cover           *common.Images `json:"cover" json:"cover"`
}

func (StudentDetail) TableName() string {
	return Student{}.TableName()
}
