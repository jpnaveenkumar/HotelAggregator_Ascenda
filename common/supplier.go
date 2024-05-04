package common

type Supplier interface {
	GetAllHotelIDs() ([]string, error)
	GetDestination(id string) (int64, error)
	GetName(id string) (string, error)
	GetLatitude(id string) (float64, error)
	GetLongitude(id string) (float64, error)
	GetAddress(id string) (string, error)
	GetCity(id string) (string, error)
	GetCountry(id string) (string, error)
	GetDescription(id string) (string, error)
	GetAmenities(id string) (map[string][]string, error)
	GetImages(id string) (map[string][]*ImageMetadata, error)
	GetBookingConditions(id string) ([]string, error)
}
