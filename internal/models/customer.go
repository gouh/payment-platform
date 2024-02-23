package models

import "github.com/google/uuid"

type Customer struct {
	ID    uuid.UUID `db:"id" json:"id"`
	Name  string    `db:"name" json:"name" binding:"required"`
	Email string    `db:"email" json:"email" binding:"required"`
}
