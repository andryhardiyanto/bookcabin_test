package model

import "backend/pkg/datatype"

type Voucher struct {
	ID           int    `json:"id" db:"id"`
	CrewName     string `json:"crewName" db:"crew_name"`
	CrewID       string `json:"crewId" db:"crew_id"`
	FlightNumber string `json:"flightNumber" db:"flight_number"`
	FlightDate   string `json:"flightDate" db:"flight_date"`
	AircraftType string `json:"aircraftType" db:"aircraft_type"`
	Seat1        string `json:"seat1" db:"seat1"`
	Seat2        string `json:"seat2" db:"seat2"`
	Seat3        string `json:"seat3" db:"seat3"`
	CreatedAt    string `json:"createdAt" db:"created_at"`
}

type RequestGenerateVoucher struct {
	Name         string        `json:"name" validate:"required"`
	ID           string        `json:"id" validate:"required"`
	FlightNumber string        `json:"flightNumber" validate:"required"`
	Date         datatype.Date `json:"date" validate:"required,date"`
	Aircraft     AircraftType  `json:"aircraft" validate:"required,oneof='ATR' 'Airbus 320' 'Boeing 737 Max'"`
}

type RequestDatabaseGenerateVoucher struct {
	CrewName     string
	CrewID       string
	FlightNumber string
	Date         datatype.Date
	AircraftType AircraftType
	Seat1        string
	Seat2        string
	Seat3        string
}

type ResponseGenerateVoucher struct {
	Seats []string `json:"seats"`
}

type RequestCheckVoucher struct {
	FlightNumber string        `json:"flightNumber" validate:"required"`
	Date         datatype.Date `json:"date" validate:"required,date"`
}

type ResponseCheckVoucher struct {
	Exists bool `json:"exists"`
}
