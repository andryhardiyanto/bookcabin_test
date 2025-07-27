package model

type AircraftType string

const (
	AircraftTypeATR    AircraftType = "ATR"
	AircraftType320    AircraftType = "Airbus 320"
	AircraftType737Max AircraftType = "Boeing 737 Max"
)

var SeatMaps = map[AircraftType][]string{
	AircraftTypeATR:    generateSeats(1, 18, []string{"A", "C", "D", "F"}),
	AircraftType320:    generateSeats(1, 32, []string{"A", "B", "C", "D", "E", "F"}),
	AircraftType737Max: generateSeats(1, 32, []string{"A", "B", "C", "D", "E", "F"}),
}
