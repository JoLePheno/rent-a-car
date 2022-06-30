package postgres

import (
	"project/internal/port"

	"github.com/go-pg/pg"
)

var _ port.RentalStore = (*RentalStore)(nil)

type RentalStore struct {
	db *pg.DB
}

func NewRentalStore(db *pg.DB) *RentalStore {
	return &RentalStore{db: db}
}
