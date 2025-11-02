package utils

import (
	"errors"
	"math/rand/v2"
	"strconv"
	"strings"
)

func Join(slice []string, sep string) string {
	result := ""
	for i, s := range slice {
		if i > 0 {
			result += sep
		}
		result += s
	}
	return result
}
func SeatListForAircraft(aircraft string) ([]string, error) {
	fn, ok := seatLayouts[aircraft]
	if !ok {
		return nil, errors.New("unknown aircraft type")
	}
	return fn(), nil
}

var seatLayouts = map[string]func() []string{
	"ATR": func() []string {
		return buildSeats(1, 18, []string{"A", "C", "D", "F"})
	},
	"Airbus 320": func() []string {
		return buildSeats(1, 32, []string{"A", "B", "C", "D", "E", "F"})
	},
	"Boeing 737 Max": func() []string {
		return buildSeats(1, 32, []string{"A", "B", "C", "D", "E", "F"})
	},
}

func buildSeats(startRow, endRow int, letters []string) []string {
	out := make([]string, 0, (endRow-startRow+1)*len(letters))
	for r := startRow; r <= endRow; r++ {
		for _, l := range letters {
			out = append(out, formatSeat(r, l))
		}
	}
	return out
}

func formatSeat(row int, letter string) string {
	return strings.TrimSpace(strings.Join([]string{itoa(row), letter}, ""))
}

func itoa(i int) string { return strconv.Itoa(i) }

func PickRandomSeats(list []string, n int) ([]string, error) {
	if n <= 0 {
		return nil, errors.New("value must be > 0")
	}
	if len(list) < n {
		return nil, errors.New("not enough seats")
	}
	s := make([]string, len(list))
	copy(s, list)
	rand.Shuffle(len(s), func(i, j int) { s[i], s[j] = s[j], s[i] })
	return s[:n], nil
}
