package patagonia

import (
	"ascenda/common"
	"errors"
	"fmt"
	"strings"
)

func (d DataSource) getHotel(id string) (*hotel, error) {
	hotelData, ok := d.dataSource[id]
	if !ok {
		return nil, errors.New("no hotel found")
	}
	return hotelData, nil
}

func (d DataSource) GetAllHotelIDs() ([]string, error) {
	return d.hotelIDs, nil
}

func (d DataSource) GetDestination(id string) (int64, error) {
	hotelData, err := d.getHotel(id)
	if err != nil {
		return -1, err
	}
	if hotelData.Destination == 0 {
		return 0, fmt.Errorf("invalid destinationID: %v", hotelData.Destination)
	}
	return hotelData.Destination, nil
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
	if hotelData.Latitude == 0 {
		return 0, fmt.Errorf("invalid Latitude: %v", hotelData.Latitude)
	}
	return hotelData.Latitude, nil
}

func (d DataSource) GetLongitude(id string) (float64, error) {
	hotelData, err := d.getHotel(id)
	if err != nil {
		return 0, err
	}
	if hotelData.Longitude == 0 {
		return 0, fmt.Errorf("invalid Longitude: %v", hotelData.Longitude)
	}
	return hotelData.Longitude, nil
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
	return "", errors.New("no city found")
}

func (d DataSource) GetCountry(id string) (string, error) {
	return "", errors.New("no country found")
}

func (d DataSource) GetDescription(id string) (string, error) {
	hotelData, err := d.getHotel(id)
	if err != nil {
		return "", err
	}
	information := strings.TrimSpace(hotelData.Information)
	if len(information) == 0 {
		return "", fmt.Errorf("invalid description: %v", information)
	}
	return information, nil
}

func (d DataSource) GetAmenities(id string) (map[string][]string, error) {
	hotelData, err := d.getHotel(id)
	if err != nil {
		return nil, err
	}
	facilities := hotelData.Amenities
	if len(facilities) == 0 {
		return nil, fmt.Errorf("invalid facilities: %v", facilities)
	}
	amenities := map[string][]string{
		"general": facilities,
	}
	return amenities, nil
}

func (d DataSource) GetImages(id string) (map[string][]*common.ImageMetadata, error) {
	hotelData, err := d.getHotel(id)
	if err != nil {
		return nil, err
	}
	images := hotelData.Images
	if len(images) == 0 {
		return nil, fmt.Errorf("invalid images: %v", images)
	}
	transformedImages := map[string][]*common.ImageMetadata{}
	for key, imagesArr := range images {
		var imagesMetaData []*common.ImageMetadata
		for _, imageDescription := range imagesArr {
			imagesMetaData = append(imagesMetaData, &common.ImageMetadata{
				Link:        imageDescription.URL,
				Description: imageDescription.Description,
			})
		}
		transformedImages[key] = imagesMetaData
	}
	return transformedImages, nil
}

func (d DataSource) GetBookingConditions(id string) ([]string, error) {
	return nil, errors.New("no booking conditions found")
}
