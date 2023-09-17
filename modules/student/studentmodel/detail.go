package studentmodel

import (
	"studyGoApp/common"
)

type StudentDetail struct {
	common.SQLModel `json: ", inline"`
	Password        string         `db:"password" json:"-"`
	FbId            string         `db:"fb_id" json:"-"`
	GgId            string         `db:"gg_id" json:"-"`
	Salt            string         `db:"salt" json:"-"`
	Role            string         `db:"role" json:"-"`
	StudentID       string         `json:"studentID" db:"studentID"`
	Name            string         `json:"name" db:"name"`
	Birthday        string         `json:"birthday" db:"birthday"`
	Avatar          *common.Image  `json:"avatar" json:"avatar"`
	Cover           *common.Images `json:"cover" json:"cover"`
}

func (StudentDetail) TableName() string {
	return Student{}.TableName()
}

func (data *StudentDetail) Mask(isAdminOrOwner bool) {
	data.GenUID(common.DbTypeStudent)
}
