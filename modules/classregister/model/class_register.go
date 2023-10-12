package classregistermodel

import (
	"studyGoApp/common"
)

type Register struct {
	StudentId int    `db:"student_id" json:"student_id"`
	ClassId   int    `db:"class_id" json:"class_id"`
	CreatedAt string `db:"created_at" json:"created_at"`

	Students *common.SimpleStudent `json:"students"`
}

func (r *Register) GetStudentId() int {
	return r.StudentId
}

func (r *Register) GetClassId() int {
	return r.ClassId
}

func (r Register) TableName() string {
	return "class_registers"
}

const EntityName = "class_registers"

func ErrorIsRegistered(err error) *common.AppError {
	return common.NewCustomError(
		err,
		"You already have a registered this class!",
		"ClassRegistered",
	)
}

func ErrorCannotCancelRegistration(err error) *common.AppError {
	return common.NewCustomError(
		err,
		"You haven't a registered this class!",
		"ClassRegister",
	)
}
