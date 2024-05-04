package common

type Location struct {
	Latitude  float64 `json:"lat"`
	Longitude float64 `json:"lng"`
	Address   string  `json:"address"`
	City      string  `json:"city"`
	Country   string  `json:"country"`
}

type ImageMetadata struct {
	Link        string `json:"link"`
	Description string `json:"description"`
}

type Hotel struct {
	ID                string                      `json:"id"`
	DestinationID     int64                       `json:"destination_id"`
	Name              string                      `json:"name"`
	Location          *Location                   `json:"location"`
	Description       string                      `json:"description"`
	Amenities         map[string][]string         `json:"amenities"`
	Images            map[string][]*ImageMetadata `json:"images"`
	BookingConditions []string                    `json:"booking_conditions"`
}
