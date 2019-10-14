package station

type ClimateStation struct {
	Name          string              `json:"name"`
	AvailableData map[string][]string `json:"availableData"`
}

func NewClimateStation() *ClimateStation {
	return &ClimateStation{
		Name:          "",
		AvailableData: make(map[string][]string),
	}
}
