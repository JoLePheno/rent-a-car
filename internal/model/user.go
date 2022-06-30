package model

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID          uuid.UUID
	Name        string
	Email       string
	PhoneNumber string
	CreatedAt   time.Time
	UpdateAt    time.Time
}
