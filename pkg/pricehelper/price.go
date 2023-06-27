package pricehelper

import (
	"math"

	"github.com/aizeNR/user-balance-service/pkg/exmath"
)

const pennyCoef = 100
const precision = 2

func RublesToPenny(r float64) uint64 {
	return uint64(math.Round(exmath.RoundFloat(r, precision) * pennyCoef))
}

func PennyToRubles(p uint64) float64 {
	return exmath.RoundFloat(float64(p)/pennyCoef, precision)
}
