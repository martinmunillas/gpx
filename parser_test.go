package gpxparser

import (
	"strings"
	"testing"
)

func TestParseValidGPX(t *testing.T) {
	gpxData := `
	<gpx>
		<trk>
			<trkseg>
				<trkpt lat="45.0" lon="9.0">
					<ele>100.0</ele>
				</trkpt>
				<trkpt lat="46.0" lon="10.0">
					<ele>200.0</ele>
				</trkpt>
			</trkseg>
		</trk>
	</gpx>`

	r := strings.NewReader(gpxData)
	gpx, err := Parse(r)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if len(gpx.Tracks) != 1 {
		t.Errorf("Expected 1 track, got %d", len(gpx.Tracks))
	}

	if len(gpx.Tracks[0].Segments) != 1 {
		t.Errorf("Expected 1 segment, got %d", len(gpx.Tracks[0].Segments))
	}

	if len(gpx.Tracks[0].Segments[0].TrackPoints) != 2 {
		t.Errorf("Expected 2 track points, got %d", len(gpx.Tracks[0].Segments[0].TrackPoints))
	}

	pt := gpx.Tracks[0].Segments[0].TrackPoints[0]
	if pt.Lat != 45.0 || pt.Lon != 9.0 || pt.Ele != 100.0 {
		t.Errorf("Expected track point (45.0, 9.0, 100.0), got (%v, %v, %v)", pt.Lat, pt.Lon, pt.Ele)
	}
}

func TestParseMalformedGPX(t *testing.T) {
	gpxData := `<gpx><trk><trkseg><trkpt lat="45.0" lon="9.0"><ele>100.0</ele></trkpt></trkseg></gpx`

	r := strings.NewReader(gpxData)
	_, err := Parse(r)
	if err == nil {
		t.Fatal("Expected error, got nil")
	}
}

func TestParseEmptyInput(t *testing.T) {
	r := strings.NewReader("")
	_, err := Parse(r)
	if err == nil {
		t.Fatal("Expected error, got nil")
	}
}

func TestParseComplexGPX(t *testing.T) {
	gpxData := `
	<gpx>
		<trk>
			<trkseg>
				<trkpt lat="45.0" lon="9.0">
					<ele>100.0</ele>
				</trkpt>
			</trkseg>
			<trkseg>
				<trkpt lat="46.0" lon="10.0">
					<ele>200.0</ele>
				</trkpt>
			</trkseg>
		</trk>
		<trk>
			<trkseg>
				<trkpt lat="47.0" lon="11.0">
					<ele>300.0</ele>
				</trkpt>
			</trkseg>
		</trk>
	</gpx>`

	r := strings.NewReader(gpxData)
	gpx, err := Parse(r)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if len(gpx.Tracks) != 2 {
		t.Errorf("Expected 2 tracks, got %d", len(gpx.Tracks))
	}

	if len(gpx.Tracks[0].Segments) != 2 {
		t.Errorf("Expected 2 segments in the first track, got %d", len(gpx.Tracks[0].Segments))
	}

	if len(gpx.Tracks[1].Segments) != 1 {
		t.Errorf("Expected 1 segment in the second track, got %d", len(gpx.Tracks[1].Segments))
	}

	pt := gpx.Tracks[1].Segments[0].TrackPoints[0]
	if pt.Lat != 47.0 || pt.Lon != 11.0 || pt.Ele != 300.0 {
		t.Errorf("Expected track point (47.0, 11.0, 300.0), got (%v, %v, %v)", pt.Lat, pt.Lon, pt.Ele)
	}
}
