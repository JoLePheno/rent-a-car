package model

import (
	"time"

	"github.com/google/uuid"
)

type Car struct {
	ID         uuid.UUID
	RenterUUID uuid.UUID

	Model        string
	Brand        string
	Availability bool

	RentAt time.Time
}
