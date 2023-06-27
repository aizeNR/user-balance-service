package exmath

import "math"

func RoundFloat(val float64, precision uint) float64 {
	//nolint:gomnd //коэф для определения разряда
	ratio := math.Pow(10, float64(precision))

	return math.Round(val*ratio) / ratio
}
