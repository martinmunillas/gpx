package gpx

type ElevationPoint struct {
	Distance  float64 `json:"distance"`
	Elevation float64 `json:"elevation"`
}

func (gpx *GPX) ToElevationData() []ElevationPoint {
	gpx.WithDistances()

	length := 0
	for _, track := range gpx.Tracks {
		for _, segment := range track.Segments {
			length += len(segment.TrackPoints)
		}
	}
	s := make([]ElevationPoint, 0, length)

	for _, track := range gpx.Tracks {
		for _, segment := range track.Segments {
			for _, point := range segment.TrackPoints {
				s = append(s, ElevationPoint{
					Elevation: point.Elevation,
					Distance:  point.RunningDistance,
				})
			}
		}
	}

	return s
}
