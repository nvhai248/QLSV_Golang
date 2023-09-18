package common

const (
	DbTypeStudent = 1
	DbTypeClass   = 2
)

const CurrentStudent = "student"

type Requester interface {
	GetId() int
	GetStudentId() string
	GetRole() string
}
