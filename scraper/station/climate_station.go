package station

type ClimateStation struct {
	Name string
	AvailableData map[string][]string
}

func NewClimateStation() *ClimateStation {
	return &ClimateStation{
		Name: "",
		AvailableData: make(map[string][]string),
	}
}
