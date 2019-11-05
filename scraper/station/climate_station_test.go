package station

import "testing"

func TestNewClimateStationName(t *testing.T) {
	climateStation := NewClimateStation()
	if climateStation.Name != "" {
		t.Fatalf("[TestNewClimateStationName] Climate Station Name: \"%s\", expected: \"\"", climateStation.Name)
	}
}

func TestNewClimateStationDataNullCheck(t *testing.T) {
	climateStation := NewClimateStation()
	if climateStation.AvailableData == nil {
		t.Fatalf("[TestNewClimateStationDataNullCheck] Climate Station Data: nil, expected: instaniated")
	}
}

func TestNewClimateStationDataLengthCheck(t *testing.T) {
	climateStation := NewClimateStation()
	if climateStation.AvailableData == nil {
		t.Fatalf("[TestNewClimateStationDataLengthCheck] Climate Station Data: nil, expected: instaniated")
	}

	if len(climateStation.AvailableData) != 0 {
		t.Fatalf("[TestNewClimateStationDataLengthCheck] Climate Station Data length: %d, expected: 0", len(climateStation.AvailableData))
	}
}
