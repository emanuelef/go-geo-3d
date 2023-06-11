package position

import (
	"math"
)

// reference: https://en.wikipedia.org/wiki/World_Geodetic_System#1984_version
const (
	WGS84_a = 6378137.0
	WGS84_b = 6356752.314245
	WGS84_f = 1 / 298.257223563 // WGS-84 ellipsiod
)

// Coord is a Lat Long struct.
type Coord2D struct {
	Lat float64 `json:"lat"`
	Lon float64 `json:"lon"`
}

type Coord3D struct {
	Coord2D
	Alt float64 `json:"alt"`
}

type Coord4D struct {
	Coord3D
	Timestamp int64 `json:"timestamp"`
}

func NewCoord3d(lat, lon, alt float64) Coord3D {
	return Coord3D{
		Coord2D: Coord2D{Lat: lat, Lon: lon},
		Alt:     alt,
	}
}

func NewCoord4d(lat, lon, alt float64, timestamp int64) Coord4D {
	return Coord4D{
		Coord3D:   NewCoord3d(lat, lon, alt),
		Timestamp: timestamp,
	}
}

func hsin(theta float64) float64 {
	return math.Pow(math.Sin(theta/2), 2)
}

func HaversineDistance(p1, p2 Coord2D) float64 {
	var aLat, aLon, bLat, bLon float64

	piRad := math.Pi / 180
	aLat = p1.Lat * piRad
	aLon = p1.Lon * piRad
	bLat = p2.Lat * piRad
	bLon = p2.Lon * piRad

	h := hsin(bLat-aLat) + math.Cos(aLat)*math.Cos(bLat)*hsin(bLon-aLon)

	meters := 2 * WGS84_a * math.Asin(math.Sqrt(h))
	return meters
}

func EuclideanDistance(a, b float64) float64 {
	return (math.Sqrt(math.Pow(a, 2) + math.Pow(b, 2)))
}

func Distance3D(a, b Coord3D) float64 {
	distance2d := HaversineDistance(a.Coord2D, b.Coord2D)
	return EuclideanDistance(distance2d, a.Alt-b.Alt)
}

// http://mathworld.wolfram.com/Point-LineDistance3-Dimensional.html
// TODO: Looks like it's returning on the projected line
func MinDistancePointToLine3D(startPosition, endPosition, posA Coord3D) float64 {
	AB := Distance3D(startPosition, endPosition)
	BC := Distance3D(endPosition, posA)
	AC := Distance3D(posA, startPosition)

	p := (AB + BC + AC) / 2 // half-perimeter
	Area := math.Sqrt(p * (p - AB) * (p - BC) * (p - AC))

	return (2 * Area) / AB
}

// ported from Python: https://stackoverflow.com/a/69604627/1077888
func LatLonAltToXYZWgs84(pos Coord3D) Point3D {
	a := WGS84_a // radius a of earth in meters cfr WGS84
	b := WGS84_b // radius b of earth in meters cfr WGS84
	e2 := 1 - (math.Pow(b, 2) / math.Pow(a, 2))
	latr := pos.Lat / 90 * 0.5 * math.Pi // latitude in radians
	lonr := pos.Lon / 180 * math.Pi      // longitude in radians
	Nphi := a / math.Sqrt(1-e2*math.Pow(math.Sin(latr), 2))
	x := (Nphi + pos.Alt) * math.Cos(latr) * math.Cos(lonr)
	y := (Nphi + pos.Alt) * math.Cos(latr) * math.Sin(lonr)
	z := (math.Pow(b, 2)/math.Pow(a, 2)*Nphi + pos.Alt) * math.Sin(latr)
	return Point3D{
		X: x,
		Y: y,
		Z: z,
	}
}

func RadiansToDegrees(radians float64) float64 {
	return radians * (180.0 / math.Pi)
}

// ported from Python: https://stackoverflow.com/a/67078296/1077888
func XYZWgs84ToLatLon(point Point3D) Coord3D {
	a := WGS84_a // in meters
	b := WGS84_b // in meters

	f := (a - b) / a

	e_sq := f * (2 - f)
	eps := e_sq / (1.0 - e_sq)

	p := math.Sqrt(math.Pow(point.X, 2) + math.Pow(point.Y, 2))
	q := math.Atan2((point.Z * a), (p * b))

	sin_q := math.Sin(q)
	cos_q := math.Cos(q)

	sin_q_3 := sin_q * sin_q * sin_q
	cos_q_3 := cos_q * cos_q * cos_q

	phi := math.Atan2((point.Z + eps*b*sin_q_3), (p - e_sq*a*cos_q_3))
	lam := math.Atan2(point.Y, point.X)

	v := a / math.Sqrt(1.0-e_sq*math.Sin(phi)*math.Sin(phi))
	h := (p / math.Cos(phi)) - v

	lat := RadiansToDegrees(phi)
	lon := RadiansToDegrees(lam)

	return Coord3D{Coord2D: Coord2D{lat, lon}, Alt: h}
}

// https://gamedev.stackexchange.com/a/72529
func (p Coord3D) ClosestPointOnLine(a, b Coord3D) (Coord3D, error) {
	A := LatLonAltToXYZWgs84(a)
	B := LatLonAltToXYZWgs84(b)
	P := LatLonAltToXYZWgs84(p)

	AP := P.Sub(A)
	AB := B.Sub(A)

	scalar := AP.Dot(AB) / AB.Dot(AB)

	if scalar < 0 || scalar > 1 {
		// The projected point is on the extended line and not lying within the segment
		if P.Distance(A) < P.Distance(B) {
			return a, nil
		} else {
			return b, nil
		}
		// return Coord3D{}, errors.New("Point not on the segment")
	}

	res := A.Add(AB.MultiplyByScalar(scalar))

	return XYZWgs84ToLatLon(res), nil
}

func (p Coord3D) ClosestPointOnLineWithTimestamp(a, b Coord4D) (Coord4D, error) {
	A := a.Coord3D
	B := b.Coord3D

	res, err := p.ClosestPointOnLine(A, B)
	if err != nil {
		return Coord4D{}, err
	}

	distanceStartEnd := Distance3D(A, B)
	distanceStartMin := Distance3D(A, p)

	// if outside segment will project the timestamp
	timestampMin := int64(float64(a.Timestamp) + (distanceStartMin/distanceStartEnd)*(float64(b.Timestamp)-float64(a.Timestamp)))

	if res == A {
		timestampMin = a.Timestamp
	}

	if res == B {
		timestampMin = b.Timestamp
	}

	return Coord4D{Coord3D: res, Timestamp: timestampMin}, nil
}
