package model

import "time"

type Task struct {
	Id          int64      `json:"id" db:"id"`
	UserId      int64      `json:"user_id" db:"user_id"`
	Title       string     `json:"title" db:"title" binding:"required"`
	Description time.Time  `json:"description" db:"description" binding:"required"`
	Status      string     `json:"status" db:"status"`
	CreatedAt   time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at" db:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at" db:"deleted_at"`
}
