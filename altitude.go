package gpx

func (gpx *GPX) ToAltitudeSlice() []float64 {
	length := 0
	for _, track := range gpx.Tracks {
		for _, segment := range track.Segments {
			length += len(segment.TrackPoints)
		}
	}
	s := make([]float64, 0, length)

	for _, track := range gpx.Tracks {
		for _, segment := range track.Segments {
			for _, point := range segment.TrackPoints {
				s = append(s, point.Elevation)
			}
		}
	}

	return s
}
