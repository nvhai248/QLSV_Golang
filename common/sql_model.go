package common

type SQLModel struct {
	Id        int    `json:"-" db:"id"`
	FakeID    *UID   `json:"id" db:"-"`
	Status    int    `json:"status" db:"status"`
	CreatedAt string `json:"created_at" db:"created_at"`
	UpdatedAt string `json:"updated_at" db:"updated_at"`
}

func (s *SQLModel) GenUID(dbType int) {
	uid := NewUID(uint32(s.Id), dbType, 1)

	s.FakeID = &uid
}
