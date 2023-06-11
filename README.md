# go-geo-3d

[![Linter](https://github.com/emanuelef/go-geo-3d/actions/workflows/linter.yml/badge.svg)](https://github.com/emanuelef/go-geo-3d/actions/workflows/linter.yml)
[![Test](https://github.com/emanuelef/go-geo-3d/actions/workflows/test.yml/badge.svg)](https://github.com/emanuelef/go-geo-3d/actions/workflows/test.yml)
![Coverage](https://raw.githubusercontent.com/emanuelef/go-geo-3d/badges/.badges/main/coverage.svg)


----

## Installation
```bash
go get github.com/emanuelef/go-geo-3d
```

## Example

```go
package main

import (
	"fmt"

	geo "github.com/emanuelef/go-geo-3d"
)

func main() {
	// Coordinates are in degrees and altitude in metres
	start := geo.NewCoord3d(51.39674, -0.36148, 1104.9)
	end := geo.NewCoord3d(51.38463, -0.36819, 1219.2)

	// Distance in metres between two 3D coordinates
	distance := geo.Distance3D(start, end)

	fmt.Printf("Distance 3D line from start to end: %.3fm\n", distance)

	posA := geo.NewCoord3d(51.3909, -0.364, 15)
	// Minimum distance in metres from one 3D point to a project line in 3D coordinates
	minPoint, _ := posA.ClosestPointOnLine(start, end)
    // Lat: 51.39181 Lon: -0.36421 Alt: 1151.37514
    
	distance = geo.Distance3D(posA, minPoint)

	fmt.Printf("Distance from one point to a line: %.3fm\n", distance)
}
```