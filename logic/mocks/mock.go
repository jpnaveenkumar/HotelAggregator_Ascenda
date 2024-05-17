package mocks

import (
	"ascenda/common"
	"errors"
)

type MockSupplier struct{}

func (m MockSupplier) GetAllHotelIDs() ([]string, error) {
	return nil, errors.New("error getting hotelIDs")
}

func (m MockSupplier) GetDestination(id string) (int64, error) {
	//TODO implement me
	panic("implement me")
}

func (m MockSupplier) GetName(id string) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (m MockSupplier) GetLatitude(id string) (float64, error) {
	//TODO implement me
	panic("implement me")
}

func (m MockSupplier) GetLongitude(id string) (float64, error) {
	//TODO implement me
	panic("implement me")
}

func (m MockSupplier) GetAddress(id string) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (m MockSupplier) GetCity(id string) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (m MockSupplier) GetCountry(id string) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (m MockSupplier) GetDescription(id string) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (m MockSupplier) GetAmenities(id string) (map[string][]string, error) {
	//TODO implement me
	panic("implement me")
}

func (m MockSupplier) GetImages(id string) (map[string][]*common.ImageMetadata, error) {
	//TODO implement me
	panic("implement me")
}

func (m MockSupplier) GetBookingConditions(id string) ([]string, error) {
	//TODO implement me
	panic("implement me")
}
