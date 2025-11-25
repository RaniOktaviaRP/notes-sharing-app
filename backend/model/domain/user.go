package domain

import (
	"time"
	"github.com/google/uuid"
)

type User struct {
	Id           uuid.UUID  `db:"id"`
	Email        string     `db:"email"`
	Name         string     `db:"name"`
	PasswordHash string     `db:"password_hash"`
	CreatedAt    time.Time  `db:"created_at"`
	UpdatedAt    time.Time  `db:"updated_at"`
	DeletedAt    *time.Time `db:"deleted_at"`
}
