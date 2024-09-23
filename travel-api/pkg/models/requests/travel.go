package request

import (
	"time"
	models "travel-api/pkg/models/travel"

	"github.com/google/uuid"
)

type TravelRequest struct {
	Name        string    `json:"name" binding:"required"`
	Destination string    `json:"destination" binding:"required"`
	Price       float64   `json:"price" binding:"required"`
	Seats       uint      `json:"seats" binding:"required"`
	Departure   time.Time `json:"departure" binding:"required"`
}

type UpdateSeatsRequest struct {
	Seats  []uuid.UUID       `json:"seats" binding:"required"`
	Status models.SeatStatus `json:"status" binding:"required"`
}
