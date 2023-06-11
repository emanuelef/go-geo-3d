package position

import (
	"math"
)

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
