package acme

import (
	"ascenda/common"
	"errors"
	"fmt"
	"strings"
)

func (d DataSource) GetAllHotelIDs() ([]string, error) {
	return d.hotelIDs, nil
}

func (d DataSource) getHotel(id string) (*hotel, error) {
	hotelData, ok := d.dataSource[id]
	if !ok {
		return nil, errors.New("no hotel found")
	}
	return hotelData, nil
}

func (d DataSource) GetDestination(id string) (int64, error) {
	hotelData, err := d.getHotel(id)
	if err != nil {
		return -1, err
	}
	if hotelData.DestinationId == 0 {
		return 0, fmt.Errorf("invalid destinationID: %v", hotelData.DestinationId)
	}
	return hotelData.DestinationId, nil
}

func (d DataSource) GetName(id string) (string, error) {
	hotelData, err := d.getHotel(id)
	if err != nil {
		return "", err
	}
	name := strings.TrimSpace(hotelData.Name)
	if len(name) == 0 {
		return "", fmt.Errorf("invalid hotel name: %v", name)
	}
	return name, nil
}

func (d DataSource) GetLatitude(id string) (float64, error) {
	hotelData, err := d.getHotel(id)
	if err != nil {
		return 0, err
	}
	if latitude, ok := hotelData.Latitude.(float64); ok {
		return latitude, nil
	}
	return 0, fmt.Errorf("invalid latitude: %v", hotelData.Latitude)
}

func (d DataSource) GetLongitude(id string) (float64, error) {
	hotelData, err := d.getHotel(id)
	if err != nil {
		return 0, err
	}
	if longitude, ok := hotelData.Longitude.(float64); ok {
		return longitude, nil
	}
	return 0, fmt.Errorf("invalid longitude: %v", hotelData.Longitude)
}

func (d DataSource) GetAddress(id string) (string, error) {
	hotelData, err := d.getHotel(id)
	if err != nil {
		return "", err
	}
	address := strings.TrimSpace(hotelData.Address)
	if len(address) == 0 {
		return "", fmt.Errorf("invalid address: %v", address)
	}
	return address, nil
}

func (d DataSource) GetCity(id string) (string, error) {
	hotelData, err := d.getHotel(id)
	if err != nil {
		return "", err
	}
	city := strings.TrimSpace(hotelData.City)
	if len(city) == 0 {
		return "", fmt.Errorf("invalid city: %v", city)
	}
	return city, nil
}

func (d DataSource) GetCountry(id string) (string, error) {
	hotelData, err := d.getHotel(id)
	if err != nil {
		return "", err
	}
	country := strings.TrimSpace(hotelData.Country)
	if len(country) == 0 {
		return "", fmt.Errorf("invalid country: %v", country)
	}
	return country, nil
}

func (d DataSource) GetDescription(id string) (string, error) {
	hotelData, err := d.getHotel(id)
	if err != nil {
		return "", err
	}
	description := strings.TrimSpace(hotelData.Description)
	if len(description) == 0 {
		return "", fmt.Errorf("invalid description: %v", description)
	}
	return description, nil
}

func (d DataSource) GetAmenities(id string) (map[string][]string, error) {
	hotelData, err := d.getHotel(id)
	if err != nil {
		return nil, err
	}
	facilities := hotelData.Facilities
	if len(facilities) == 0 {
		return nil, fmt.Errorf("invalid facilities: %v", facilities)
	}
	amenities := map[string][]string{
		"general": facilities,
	}
	return amenities, nil
}

func (d DataSource) GetImages(id string) (map[string][]*common.ImageMetadata, error) {
	return nil, errors.New("no images found")
}

func (d DataSource) GetBookingConditions(id string) ([]string, error) {
	return nil, errors.New("no booking conditions found")
}
