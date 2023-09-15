package studentmodel

import (
	"errors"
	"strings"
	"studyGoApp/common"
)

const EntityName = "Student"

type Student struct {
	common.SQLModel `json:", inline"`

	StudentID  string         `db:"studentID" json:"studentID"`
	Birthday   string         `db:"birthday" json:"birthday"`
	Name       string         `db:"name" json:"name"`
	Avatar     *common.Image  `json:"avatar" json:"avatar"`
	Cover      *common.Images `json:"cover" json:"cover"`
	ClassCount int            `json:"class_count" json:"-"`
}

func (Student) TableName() string {
	return "student"
}

type StudentUpdate struct {
	StudentID *string        `db:"studentID" json:"studentID"`
	Birthday  *string        `db:"birthday" json:"birthday"`
	Name      *string        `db:"name" json:"name"`
	Avatar    *common.Image  `json:"avatar" json:"avatar"`
	Cover     *common.Images `json:"cover" json:"cover"`
}

func (StudentUpdate) TableName() string {
	return Student{}.TableName()
}

type StudentCreate struct {
	common.SQLModel `json:", inline"`

	StudentID string         `db:"studentID" json:"studentID"`
	Birthday  string         `db:"birthday" json:"birthday"`
	Name      string         `db:"name" json:"name"`
	Avatar    *common.Image  `json:"avatar" json:"avatar"`
	Cover     *common.Images `json:"cover" json:"cover"`
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

var (
	ErrNameCannotBeEmpty = common.NewCustomError(nil, "student name can't be blank", "StudentNameErr")
)

func (data *Student) Mask(isAdminOrOwner bool) {
	data.GenUID(common.DbTypeStudent)
}

func (data *StudentCreate) Mask(isAdminOrOwner bool) {
	data.GenUID(common.DbTypeStudent)
}
