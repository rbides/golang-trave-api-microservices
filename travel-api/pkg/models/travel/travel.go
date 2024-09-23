package models

import (
	"time"

	"github.com/google/uuid"
)

type Money float64

func (m Money) ToCents() Cents {
	return Cents(m * 100)
}

type Cents uint

func (c Cents) ToMoney() Money {
	return Money(float64(c) / 100)
}

type Travel struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Destination string    `json:"destination"`
	Price       Money     `json:"price"`
	Departure   time.Time `json:"departure"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
