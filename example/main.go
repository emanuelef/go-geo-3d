package main

import (
	"fmt"

	geo "github.com/emanuelef/go-geo-3d"
)

func main() {
	start := geo.NewCoord3d(51.39674, -0.36148, 1104.9)
	end := geo.NewCoord3d(51.38463, -0.36819, 1219.2)
	posA := geo.NewCoord3d(51.3909, -0.364, 15)

	meters := geo.MinDistancePointToLine3D(start, end, posA)
	minPoint, _ := posA.ClosestPointOnLine(start, end)
	distance := geo.Distance3D(posA, minPoint)

	fmt.Printf("Distance: %.3f m %.3f m\n", meters, distance)
}
