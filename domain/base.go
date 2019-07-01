package domain

import (
	"math/big"
	"time"
)

type Base struct {
	Id        big.Int    `json:"id"`
	Slug      string     `json:"slug"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}
