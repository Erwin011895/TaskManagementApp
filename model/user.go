package model

import "time"

type User struct {
	Id        int64      `json:"id" db:"id"`
	Name      string     `json:"name" db:"name"`
	Email     string     `json:"email" db:"email" binding:"required"`
	Password  string     `json:"password,omitempty" db:"password"`
	CreatedAt time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt time.Time  `json:"updated_at" db:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at" db:"deleted_at"`
}
