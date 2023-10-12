package subscriber

type HasStudentId interface {
	GetStudentId() int
}

type HasClassId interface {
	GetClassId() int
}
