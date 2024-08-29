package gpx

import (
	"encoding/xml"
	"fmt"
	"io"
)

type GPX struct {
	XMLName xml.Name `xml:"gpx"`
	Tracks  []Track  `xml:"trk" json:"tracks"`
}

type Track struct {
	XMLName  xml.Name       `xml:"trk"`
	Segments []TrackSegment `xml:"trkseg" json:"segments"`
}

type TrackSegment struct {
	XMLName     xml.Name     `xml:"trkseg"`
	TrackPoints []TrackPoint `xml:"trkpt" json:"points"`
}

type TrackPoint struct {
	XMLName   xml.Name `xml:"trkpt"`
	Latitude  float64  `xml:"lat,attr" json:"lat"`
	Longitude float64  `xml:"lon,attr" json:"lon"`
	Elevation float64  `xml:"ele" json:"ele"`
}

func Parse(r io.Reader) (*GPX, error) {
	var gpx GPX
	if err := xml.NewDecoder(r).Decode(&gpx); err != nil {
		return nil, fmt.Errorf("error decoding GPX: %w", err)
	}

	return &gpx, nil
}
