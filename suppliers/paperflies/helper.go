package paperflies

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
	if hotelData.DestinationID == 0 {
		return 0, fmt.Errorf("invalid destinationID: %v", hotelData.DestinationID)
	}
	return hotelData.DestinationID, nil
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
	return 0, errors.New("no latitude found")
}

func (d DataSource) GetLongitude(id string) (float64, error) {
	return 0, errors.New("no longitude found")
}

func (d DataSource) GetAddress(id string) (string, error) {
	hotelData, err := d.getHotel(id)
	if err != nil {
		return "", err
	}
	if hotelData.Location == nil {
		return "", fmt.Errorf("no address found")
	}
	address := strings.TrimSpace(hotelData.Location.Address)
	if len(address) == 0 {
		return "", fmt.Errorf("invalid address: %v", address)
	}
	return address, nil
}

func (d DataSource) GetCity(id string) (string, error) {
	return "", fmt.Errorf("no city found")
}

func (d DataSource) GetCountry(id string) (string, error) {
	hotelData, err := d.getHotel(id)
	if err != nil {
		return "", err
	}
	if hotelData.Location == nil {
		return "", fmt.Errorf("no country found")
	}
	country := strings.TrimSpace(hotelData.Location.Country)
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
	information := strings.TrimSpace(hotelData.Details)
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
	return facilities, nil
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
				Link:        imageDescription.Link,
				Description: imageDescription.Caption,
			})
		}
		transformedImages[key] = imagesMetaData
	}
	return transformedImages, nil
}

func (d DataSource) GetBookingConditions(id string) ([]string, error) {
	hotelData, err := d.getHotel(id)
	if err != nil {
		return nil, err
	}
	conditions := hotelData.BookingConditions
	if len(conditions) == 0 {
		return nil, fmt.Errorf("invalid conditions: %v", conditions)
	}
	return conditions, nil
}
