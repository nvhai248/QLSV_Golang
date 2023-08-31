package studentmodel

import (
	"errors"
	"strings"
	"studyGoApp/common"
)

const EntityName = "Student"

type Student struct {
	common.SQLModel `json:", inline"`

	StudentID string `db:"studentID" json:"studentID"`
	Birthday  string `db:"birthday" json:"birthday"`
	Name      string `db:"name" json:"name"`
}

func (Student) TableName() string {
	return "student"
}

type StudentUpdate struct {
	StudentID *string `db:"studentID" json:"studentID"`
	Birthday  *string `db:"birthday" json:"birthday"`
	Name      *string `db:"name" json:"name"`
}

func (StudentUpdate) TableName() string {
	return Student{}.TableName()
}

type StudentCreate struct {
	StudentID string `db:"studentID" json:"studentID"`
	Birthday  string `db:"birthday" json:"birthday"`
	Name      string `db:"name" json:"name"`
}

func (StudentCreate) TableName() string {
	return Student{}.TableName()
}

func (stu *StudentCreate) Validate() error {
	stu.Name = strings.TrimSpace(stu.Name)

	if len(stu.Name) == 0 {
		return errors.New("student name cannot be blank!")
	}

	return nil
}
