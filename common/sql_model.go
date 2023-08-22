package common

type SQLModel struct {
	Id        int    `json:"id" db:"id"`
	Status    int    `json:"status" db:"status"`
	CreatedAt string `json:"created_at" db:"created_at"`
	UpdatedAt string `json:"updated_at" db:"updated_at"`
}
