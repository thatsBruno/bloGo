// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package database

import (
	"time"

	"github.com/google/uuid"
)

type Feed struct {
	ID        uuid.UUID
	CreateAt  time.Time
	UpdatedAt time.Time
	Name      string
	Url       string
	UserID    uuid.UUID
}

type User struct {
	ID        uuid.UUID
	CreateAt  time.Time
	UpdatedAt time.Time
	Name      string
	ApiKey    string
}
