package position

import (
	"fmt"
	"testing"
)

func TestMin3DDistance(t *testing.T) {
	// 619.3074816643333
	start := Coord3D{Coord2D: Coord2D{51.4142, -0.3519439}, Alt: 883.92}
	end := Coord3D{Coord2D: Coord2D{51.4018637, -0.3540802}, Alt: 1021}
	posA := Coord3D{Coord2D: Coord2D{51.443807, -0.343025}, Alt: 15}

	meters := MinDistancePointToLine3D(start, end, posA)
	if int(meters) != 594 {
		t.Errorf("Computed values: %10f\n", meters)
		t.Errorf("Incorrect computation between A and B: %v\n", meters)
	}
}

func BenchmarkBigLen(b *testing.B) {
	start := Coord3D{Coord2D: Coord2D{51.4142, -0.3519439}, Alt: 883.92}
	end := Coord3D{Coord2D: Coord2D{51.4018637, -0.3540802}, Alt: 1021}
	posA := Coord3D{Coord2D: Coord2D{51.443807, -0.343025}, Alt: 15}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		MinDistancePointToLine3D(start, end, posA)
	}
}

func TestConvertToXYZ(t *testing.T) {
	p := LatLonAltToXYZWgs84(Coord3D{Coord2D: Coord2D{51.443807, -0.343025}, Alt: 15.})
	fmt.Println(p)
}

func TestConvertToLatLonAlt(t *testing.T) {
	res := XYZWgs84ToLatLon(Point3D{3983477.358112, -23849.027945, 4964479.885483})
	fmt.Println(res)
}

func TestClosestPointOnLine(t *testing.T) {
	start := Coord3D{Coord2D: Coord2D{31.4142, -0.3519439}, Alt: 883.92}
	end := Coord3D{Coord2D: Coord2D{31.4018637, -0.3540802}, Alt: 1021}
	posA := Coord3D{Coord2D: Coord2D{51.443807, -0.343025}, Alt: 15}
	res, _ := posA.ClosestPointOnLine(start, end)
	fmt.Println(res)
}

func TestMin3DDistanceNew(t *testing.T) {
	start := Coord3D{Coord2D: Coord2D{51.39674299854343, -0.361480712890625}, Alt: 1104.9}
	end := Coord3D{Coord2D: Coord2D{51.384639416710804, -0.368194580078125}, Alt: 1219.2}
	posA := Coord3D{Coord2D: Coord2D{51.3909, -0.364}, Alt: 15}

	meters := MinDistancePointToLine3D(start, end, posA)
	minPoint, _ := posA.ClosestPointOnLine(start, end)
	distance := Distance3D(posA, minPoint)

	if int(meters) != int(distance) {
		t.Errorf("Computed values: %10f\n", meters)
		t.Errorf("Incorrect computation between A and B: %v\n", meters)
	}
}

func TestMin3DDistanceH(t *testing.T) {
	start := Coord3D{Coord2D: Coord2D{51.39721, -0.504455}, Alt: 381.0}
	end := Coord3D{Coord2D: Coord2D{51.39619, -0.50091}, Alt: 388.62}
	posA := Coord3D{Coord2D: Coord2D{51.397, -0.5026}, Alt: 15}

	meters := MinDistancePointToLine3D(start, end, posA)
	minPoint, _ := posA.ClosestPointOnLine(start, end)
	distance := Distance3D(posA, minPoint)

	if int(meters) != int(distance) {
		t.Errorf("Computed values: %10f\n", meters)
		t.Errorf("Incorrect computation between A and B: %v\n", meters)
	}
}
