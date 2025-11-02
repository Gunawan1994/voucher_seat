package domain

import (
	"time"
)

type Assignments struct {
	ID           int       `gorm:"primaryKey;autoIncrement" json:"id"`
	CrewName     string    `json:"crew_name"`
	CrewID       string    `json:"crew_id"`
	FlightNumber string    `json:"flight_number"`
	FlightDate   string    `json:"flight_date"`
	AircraftType string    `json:"aircraft_type"`
	Seat1        string    `json:"seat1"`
	Seat2        string    `json:"seat2"`
	Seat3        string    `json:"seat3"`
	CreatedAt    time.Time `gorm:"<-:create" json:"created_at"`
}
