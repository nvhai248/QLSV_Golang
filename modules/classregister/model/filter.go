package classregistermodel

type Filter struct {
	StudentId int `json:"student_id" form:"student_id"`
	ClassId   int `json:"class_id" form:"class_id"`
}

func (f *Filter) Mask() {
	f.ClassId = 0
	f.StudentId = 0
}
