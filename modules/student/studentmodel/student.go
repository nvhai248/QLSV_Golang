package studentmodel

import (
	"errors"
	"strings"
	"studyGoApp/common"
)

const EntityName = "Student"

type Student struct {
	common.SQLModel `json:", inline"`

	Password   string         `db:"password" json:"-"`
	FbId       string         `db:"fb_id" json:"-"`
	GgId       string         `db:"gg_id" json:"-"`
	Salt       string         `db:"salt" json:"-"`
	Role       string         `db:"role" json:"-"`
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
	Password  *string        `db:"password" json:"password"`
	FbId      *string        `db:"fb_id" json:"fb_id"`
	GgId      *string        `db:"gg_id" json:"gg_id"`
	Salt      *string        `db:"salt" json:"salt"`
	Role      *string        `db:"role" json:"role"`
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

	Password  string         `db:"password" json:"password"`
	FbId      string         `db:"fb_id" json:"fb_id"`
	GgId      string         `db:"gg_id" json:"gg_id"`
	Salt      string         `db:"salt" json:"salt"`
	Role      string         `db:"role" json:"role"`
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

type StudentLogin struct {
	StudentId string `db:"studentID" form:"studentId" json:"studentID"`
	Password  string `db:"password" form:"password" json:"password"`
}

func (StudentLogin) TableName() string {
	return Student{}.TableName()
}
