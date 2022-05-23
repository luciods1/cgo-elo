package c_elo

import (
	// #cgo CFLAGS: -I../internal
	// #cgo LDFLAGS: -L../internal -lelo -lm
	// #include "elo.h"
	"C"
	"math"
)

type Cielo interface {
	WinProbability(eloPlayerA, eloPlayerB int) float32
	RatingDifferential(eloPlayerA, eloPlayerB int) int
}

func RatingDifferential(a, b int, growthFactor float64) int16 {
	return int16(math.Round((float64(a) - WinProbability(a, b)*growthFactor)))
}

func CRatingDifferential(a, b int, growthFactor float64) C.short {
	return C.rating_differential(C.short(a), C.short(b), C.float(growthFactor))
}

func CWinProbability(eloPlayerA, eloPlayerB int) C.float {
	return C.win_probability(C.short(eloPlayerA), C.short(eloPlayerB))
}

func WinProbability(a, b int) float64 {
	return 1 / (math.Pow(10, (-float64(a-b)/400)) + 1)
}
