package domains

import "time"

type User struct {
	ID        string    `json:"id"`
	Username  string    `json:"username" binding:"required"`
	Password  string    `json:"password" binding:"required"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
