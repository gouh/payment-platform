package models

import "github.com/google/uuid"

type Merchant struct {
	ID   uuid.UUID `db:"id"`
	Name string    `db:"name"`
}
