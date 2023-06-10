package main

import (
	"fmt"

	geo "github.com/emanuelef/go-geo-3d"
)

func main() {
	start := geo.Coord3D{
		Coord2D: geo.Coord2D{Lat: 51.39674, Lon: -0.36148},
		Alt:     1104.9,
	}
	end := geo.Coord3D{
		Coord2D: geo.Coord2D{Lat: 51.38463, Lon: -0.36819},
		Alt:     1219.2,
	}
	posA := geo.Coord3D{
		Coord2D: geo.Coord2D{Lat: 51.3909, Lon: -0.364},
		Alt:     15,
	}

	meters := geo.MinDistancePointToLine3D(start, end, posA)
	minPoint, _ := posA.ClosestPointOnLine(start, end)
	distance := geo.Distance3D(posA, minPoint)

	fmt.Printf("Distance: %.3f m %.3f m\n", meters, distance)
}
