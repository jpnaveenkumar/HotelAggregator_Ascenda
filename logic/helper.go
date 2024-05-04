package logic

import (
	"ascenda/common"
	"errors"
	"fmt"
)

func getDestination(id string) (int64, error) {
	for _, dataSourceKey := range FieldVsDataSourceOrdering[common.FieldDestinationID] {
		destID, err := keyVsDataSource[dataSourceKey].GetDestination(id)
		if err != nil {
			fmt.Printf("datasource: %v field: %v err: %v", dataSourceKey, common.FieldDestinationID, err)
			continue
		}
		return destID, nil
	}
	return -1, errors.New("no destinationID found")
}

func getName(id string) (string, error) {
	for _, dataSourceKey := range FieldVsDataSourceOrdering[common.FieldName] {
		name, err := keyVsDataSource[dataSourceKey].GetName(id)
		if err != nil {
			fmt.Printf("datasource: %v field: %v err: %v", dataSourceKey, common.FieldName, err)
			continue
		}
		return name, nil
	}
	return "", errors.New("no name found")
}

func getAmenities(id string) (map[string][]string, error) {
	for _, dataSourceKey := range FieldVsDataSourceOrdering[common.FieldAmenities] {
		amenities, err := keyVsDataSource[dataSourceKey].GetAmenities(id)
		if err != nil {
			fmt.Printf("datasource: %v field: %v err: %v", dataSourceKey, common.FieldAmenities, err)
			continue
		}
		return amenities, nil
	}
	return nil, errors.New("no amenities found")
}

func getDescription(id string) (string, error) {
	for _, dataSourceKey := range FieldVsDataSourceOrdering[common.FieldDescription] {
		description, err := keyVsDataSource[dataSourceKey].GetDescription(id)
		if err != nil {
			fmt.Printf("datasource: %v field: %v err: %v", dataSourceKey, common.FieldDescription, err)
			continue
		}
		return description, nil
	}
	return "", errors.New("no description found")
}

func getImages(id string) (map[string][]*common.ImageMetadata, error) {
	for _, dataSourceKey := range FieldVsDataSourceOrdering[common.FieldImages] {
		images, err := keyVsDataSource[dataSourceKey].GetImages(id)
		if err != nil {
			fmt.Printf("datasource: %v field: %v err: %v", dataSourceKey, common.FieldImages, err)
			continue
		}
		return images, nil
	}
	return nil, errors.New("no images found")
}

func getBookingConditions(id string) ([]string, error) {
	for _, dataSourceKey := range FieldVsDataSourceOrdering[common.FieldBookingConditions] {
		bookingConditions, err := keyVsDataSource[dataSourceKey].GetBookingConditions(id)
		if err != nil {
			fmt.Printf("datasource: %v field: %v err: %v", dataSourceKey, common.FieldBookingConditions, err)
			continue
		}
		return bookingConditions, nil
	}
	return nil, errors.New("no bookingConditions found")
}

func getLatitude(id string) (float64, error) {
	for _, dataSourceKey := range FieldVsDataSourceOrdering[common.FieldLatitude] {
		latitude, err := keyVsDataSource[dataSourceKey].GetLatitude(id)
		if err != nil {
			fmt.Printf("datasource: %v field: %v err: %v", dataSourceKey, common.FieldLatitude, err)
			continue
		}
		return latitude, nil
	}
	return -1, errors.New("no latitude found")
}

func getLongitude(id string) (float64, error) {
	for _, dataSourceKey := range FieldVsDataSourceOrdering[common.FieldLongitude] {
		longitude, err := keyVsDataSource[dataSourceKey].GetLongitude(id)
		if err != nil {
			fmt.Printf("datasource: %v field: %v err: %v", dataSourceKey, common.FieldLongitude, err)
			continue
		}
		return longitude, nil
	}
	return -1, errors.New("no longitude found")
}

func getAddress(id string) (string, error) {
	for _, dataSourceKey := range FieldVsDataSourceOrdering[common.FieldAddress] {
		address, err := keyVsDataSource[dataSourceKey].GetAddress(id)
		if err != nil {
			fmt.Printf("datasource: %v field: %v err: %v", dataSourceKey, common.FieldAddress, err)
			continue
		}
		return address, nil
	}
	return "", errors.New("no address found")
}

func getCity(id string) (string, error) {
	for _, dataSourceKey := range FieldVsDataSourceOrdering[common.FieldCity] {
		city, err := keyVsDataSource[dataSourceKey].GetCity(id)
		if err != nil {
			fmt.Printf("datasource: %v field: %v err: %v", dataSourceKey, common.FieldCity, err)
			continue
		}
		return city, nil
	}
	return "", errors.New("no city found")
}

func getCountry(id string) (string, error) {
	for _, dataSourceKey := range FieldVsDataSourceOrdering[common.FieldCountry] {
		country, err := keyVsDataSource[dataSourceKey].GetCountry(id)
		if err != nil {
			fmt.Printf("datasource: %v field: %v err: %v", dataSourceKey, common.FieldCountry, err)
			continue
		}
		return country, nil
	}
	return "", errors.New("no country found")
}
