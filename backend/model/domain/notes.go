package domain

import (
	"time"
	"github.com/google/uuid"
)

type Note struct {
	Id        uuid.UUID  `db:"id"`
	// UserId    uuid.UUID  `db:"user_id"`
	Title     string     `db:"title"`
	Content   string     `db:"content"`
	Image 	  []byte	 `db:"image"`
	CreatedAt time.Time  `db:"created_at"`
	UpdatedAt time.Time  `db:"updated_at"`
	DeletedAt *time.Time `db:"deleted_at"`
}