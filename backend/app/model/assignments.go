package model

import (
	"context"
	"time"

	"voucher_seat/app/domain"
)

// BaseofRequest
type BaseAssignments struct {
	Id           int       `gorm:"primaryKey;autoIncrement" json:"id"`
	CrewName     string    `json:"crew_name"`
	CrewId       string    `json:"crew_id"`
	FlightNumber string    `json:"flight_number"`
	FlightDate   string    `json:"flight_date"`
	AircraftType string    `json:"aircraft_type"`
	Seat1        string    `json:"seat1"`
	Seat2        string    `json:"seat2"`
	Seat3        string    `json:"seat3"`
	CreatedAt    time.Time `gorm:"<-:create" json:"created_at"`
}

func (req BaseAssignments) ToDomain(ctx context.Context) *domain.Assignments {
	return &domain.Assignments{
		ID:           req.Id,
		CrewName:     req.CrewName,
		CrewID:       req.CrewId,
		FlightNumber: req.FlightNumber,
		FlightDate:   req.FlightDate,
		AircraftType: req.AircraftType,
		Seat1:        req.Seat1,
		Seat2:        req.Seat2,
		Seat3:        req.Seat3,
	}
}

// request
type CheckFlightNumberReq struct {
	FlightNumber string `json:"flightNumber" validate:"required"`
	Date         string `json:"date" validate:"required"`
}

type GenerateSeatNumberReq struct {
	Name         string `json:"name" validate:"required"`
	ID           string `json:"id" validate:"required"`
	FlightNumber string `json:"flightNumber" validate:"required,alphanum"`
	Date         string `json:"date" validate:"required,datetime=2006-01-02"`
	Aircraft     string `json:"aircraft" validate:"required"`
}

// response
type CheckFlightNumberRes struct {
	Exists bool `json:"exists"`
}

type GenerateSeatNumberRes struct {
	Seats []string `json:"seats"`
}
