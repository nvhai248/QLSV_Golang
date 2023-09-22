package common

type SimpleStudent struct {
	SQLModel `json:", inline"`

	Role      string `db:"role" json:"-"`
	StudentID string `db:"studentID" json:"studentID"`
	Birthday  string `db:"birthday" json:"birthday"`
	Name      string `db:"name" json:"name"`
}

func (SimpleStudent) TableName() string {
	return "student"
}

func (st *SimpleStudent) Mask(isAdmin bool) {
	st.GenUID(DbTypeStudent)
}
