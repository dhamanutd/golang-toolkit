package math

import "math"

func Round(decimal float64, place float64) float64 {
	return math.Round(decimal*math.Pow(10, float64(place))) / math.Pow(10, float64(place))
}
