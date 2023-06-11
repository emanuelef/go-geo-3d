package position

import (
	"math"
)

func EuclideanDistance(a, b float64) float64 {
	return (math.Sqrt(math.Pow(a, 2) + math.Pow(b, 2)))
}

func RadiansToDegrees(radians float64) float64 {
	return radians * (180.0 / math.Pi)
}
