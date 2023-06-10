package position

type TimedPosition struct {
	Coord3D
	Timestamp int64
}

func GenerateTimedLinearSubsamples(a, b TimedPosition, subSamples int) []TimedPosition {
	var subsamplesTimedPositions []TimedPosition

	slopeLat := (b.Lat - a.Lat) / float64(subSamples)
	slopeLon := (b.Lon - a.Lon) / float64(subSamples)
	slopeAlt := (b.Alt - a.Alt) / float64(subSamples)
	slopeTimestamp := float64(b.Timestamp-a.Timestamp) / float64(subSamples)

	for i := 0; i < subSamples; i++ {
		t := TimedPosition{
			Coord3D{
				Coord2D: Coord2D{
					a.Lat + slopeLat*float64(i),
					a.Lon + slopeLon*float64(i),
				},
				Alt: a.Alt + slopeAlt*float64(i),
			},
			int64(float64(a.Timestamp) + slopeTimestamp*float64(i)),
		}

		subsamplesTimedPositions = append(subsamplesTimedPositions, t)
	}

	subsamplesTimedPositions = append(subsamplesTimedPositions, b)

	return subsamplesTimedPositions
}
