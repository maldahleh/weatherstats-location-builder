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
		t.Fatal("[TestNewClimateStationDataNullCheck] Climate Station Data: nil, expected: instantiated")
	}
}

func TestNewClimateStationDataLengthCheck(t *testing.T) {
	climateStation := NewClimateStation()
	if climateStation.AvailableData == nil {
		t.Fatal("[TestNewClimateStationDataLengthCheck] Climate Station Data: nil, expected: instantiated")
	}

	if len(climateStation.AvailableData) != 0 {
		t.Fatalf("[TestNewClimateStationDataLengthCheck] Climate Station Data length: %d, expected: 0", len(climateStation.AvailableData))
	}
}
