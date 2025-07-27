package model

import (
	"fmt"
	"math/rand"
)

func generateSeats(rows int, maxRow int, seats []string) []string {
	var seatList []string
	for i := rows; i <= maxRow; i++ {
		for _, s := range seats {
			seatList = append(seatList, fmt.Sprintf("%d%s", i, s))
		}
	}
	return seatList
}

func GenerateRandomSeatsByAircraftType(aircraftType AircraftType) (seat1, seat2, seat3 string) {
	seat1 = SeatMaps[aircraftType][rand.Intn(len(SeatMaps[aircraftType]))]
	seat2 = SeatMaps[aircraftType][rand.Intn(len(SeatMaps[aircraftType]))]
	seat3 = SeatMaps[aircraftType][rand.Intn(len(SeatMaps[aircraftType]))]
	return
}
