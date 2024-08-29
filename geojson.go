package gpx

import "encoding/json"

// GeoJSON represents the structure of a GeoJSON object.
type GeoJSON struct {
	Type     string    `json:"type"`
	Features []Feature `json:"features"`
}

// Feature represents a GeoJSON feature.
type Feature struct {
	Type       string                 `json:"type"`
	Geometry   Geometry               `json:"geometry"`
	Properties map[string]interface{} `json:"properties"`
}

// Geometry represents the geometry of a GeoJSON feature.
type Geometry struct {
	Type        string      `json:"type"`
	Coordinates [][]float64 `json:"coordinates"`
}

func (g GeoJSON) String() ([]byte, error) {
	return json.Marshal(g)
}

func (gpx *GPX) ToGeoJSON() GeoJSON {
	geoJSON := GeoJSON{
		Type:     "FeatureCollection",
		Features: []Feature{},
	}

	for _, track := range gpx.Tracks {
		for _, segment := range track.Segments {
			feature := Feature{
				Type: "Feature",
				Geometry: Geometry{
					Type:        "LineString",
					Coordinates: [][]float64{},
				},
				Properties: map[string]interface{}{}, // Add any additional properties here
			}

			for _, point := range segment.TrackPoints {
				feature.Geometry.Coordinates = append(
					feature.Geometry.Coordinates,
					[]float64{point.Longitude, point.Latitude, point.Elevation},
				)
			}

			geoJSON.Features = append(geoJSON.Features, feature)
		}
	}

	return geoJSON
}
