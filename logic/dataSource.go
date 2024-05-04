package logic

import (
	"ascenda/common"
	"ascenda/suppliers/acme"
	"ascenda/suppliers/paperflies"
	"ascenda/suppliers/patagonia"
	"fmt"
)

var (
	hotelIDsVsDataSource      = map[string][]string{}
	keyVsDataSource           = map[string]common.Supplier{}
	FieldVsDataSourceOrdering = map[string][]string{}

	HotelIDVsHotel       = map[string]*common.Hotel{}
	DestinationIDVsHotel = map[int64][]*common.Hotel{}
)

func GetHotels(hotelIDs []string) []*common.Hotel {
	var response []*common.Hotel
	for _, hotelID := range hotelIDs {
		if hotel, ok := HotelIDVsHotel[hotelID]; ok {
			response = append(response, hotel)
		}
	}
	return response
}

func GetHotelsByDestinationID(destinationID int64) []*common.Hotel {
	if hotels, ok := DestinationIDVsHotel[destinationID]; ok {
		return hotels
	}
	return nil
}

func buildDataSourceOrderingForField() {
	FieldVsDataSourceOrdering[common.FieldID] = []string{common.SupplierACME, common.SupplierPatagonia, common.SupplierPaperFlies}
	FieldVsDataSourceOrdering[common.FieldDestinationID] = []string{common.SupplierACME, common.SupplierPatagonia, common.SupplierPaperFlies}
	FieldVsDataSourceOrdering[common.FieldName] = []string{common.SupplierACME, common.SupplierPatagonia, common.SupplierPaperFlies}
	FieldVsDataSourceOrdering[common.FieldLatitude] = []string{common.SupplierACME, common.SupplierPatagonia, common.SupplierPaperFlies}
	FieldVsDataSourceOrdering[common.FieldLongitude] = []string{common.SupplierACME, common.SupplierPatagonia, common.SupplierPaperFlies}
	FieldVsDataSourceOrdering[common.FieldAddress] = []string{common.SupplierPaperFlies, common.SupplierACME, common.SupplierPatagonia}
	FieldVsDataSourceOrdering[common.FieldCity] = []string{common.SupplierACME}
	FieldVsDataSourceOrdering[common.FieldCountry] = []string{common.SupplierPaperFlies, common.SupplierACME, common.SupplierPatagonia}
	FieldVsDataSourceOrdering[common.FieldDescription] = []string{common.SupplierPaperFlies, common.SupplierACME, common.SupplierPatagonia}
	FieldVsDataSourceOrdering[common.FieldAmenities] = []string{common.SupplierPaperFlies, common.SupplierACME, common.SupplierPatagonia}
	FieldVsDataSourceOrdering[common.FieldImages] = []string{common.SupplierPatagonia, common.SupplierPaperFlies}
	FieldVsDataSourceOrdering[common.FieldBookingConditions] = []string{common.SupplierPaperFlies}
}

func buildDataSource() {
	acmeSource, err := acme.Init()
	if err != nil {
		fmt.Printf("failed to initialize datasource : %v", common.SupplierACME)
	} else {
		keyVsDataSource[common.SupplierACME] = acmeSource
	}

	paperFliesSource, err := paperflies.Init()
	if err != nil {
		fmt.Printf("failed to initialize datasource : %v", common.SupplierPaperFlies)
	} else {
		keyVsDataSource[common.SupplierPaperFlies] = paperFliesSource
	}

	patagoniaSource, err := patagonia.Init()
	if err != nil {
		fmt.Printf("failed to initialize datasource : %v", common.SupplierPatagonia)
	} else {
		keyVsDataSource[common.SupplierPatagonia] = patagoniaSource
	}
}

func buildDataset() {
	for key, dataSource := range keyVsDataSource {
		hotelIDS, err := dataSource.GetAllHotelIDs()
		if err != nil {
			fmt.Printf("failed to fetch hotelIDs from datasource=%v with err=%v", key, err)
		}
		for _, hotelID := range hotelIDS {
			if _, ok := hotelIDsVsDataSource[hotelID]; !ok {
				hotelIDsVsDataSource[hotelID] = []string{key}
			} else {
				hotelIDsVsDataSource[hotelID] = append(hotelIDsVsDataSource[hotelID], key)
			}
		}
	}

	for hotelID, _ := range hotelIDsVsDataSource {
		hotel := &common.Hotel{}
		hotel.ID = hotelID
		HotelIDVsHotel[hotelID] = hotel

		destinationID, err := getDestination(hotelID)
		if err == nil {
			hotel.DestinationID = destinationID
			DestinationIDVsHotel[destinationID] = append(DestinationIDVsHotel[destinationID], hotel)
		} else {
			fmt.Printf("failed to fetch destinationID for hotelID=%v with err=%v", hotelID, err)
		}

		name, err := getName(hotelID)
		if err != nil {
			fmt.Printf("failed to fetch name for hotelID=%v with err=%v", hotelID, err)
		} else {
			hotel.Name = name
		}

		amenities, err := getAmenities(hotelID)
		if err != nil {
			fmt.Printf("failed to fetch amenities for hotelID=%v with err=%v", hotelID, err)
		} else {
			hotel.Amenities = amenities
		}

		description, err := getDescription(hotelID)
		if err != nil {
			fmt.Printf("failed to fetch description for hotelID=%v with err=%v", hotelID, err)
		} else {
			hotel.Description = description
		}

		images, err := getImages(hotelID)
		if err != nil {
			fmt.Printf("failed to fetch images for hotelID=%v with err=%v", hotelID, err)
		} else {
			hotel.Images = images
		}

		bookingConditions, err := getBookingConditions(hotelID)
		if err != nil {
			fmt.Printf("failed to fetch bookingConditions for hotelID=%v with err=%v", hotelID, err)
		} else {
			hotel.BookingConditions = bookingConditions
		}

		location := &common.Location{}
		latitude, err := getLatitude(hotelID)
		if err != nil {
			fmt.Printf("failed to fetch latitude for hotelID=%v with err=%v", hotelID, err)
		} else {
			location.Latitude = latitude
		}

		longitude, err := getLongitude(hotelID)
		if err != nil {
			fmt.Printf("failed to fetch longitude for hotelID=%v with err=%v", hotelID, err)
		} else {
			location.Longitude = longitude
		}

		address, err := getAddress(hotelID)
		if err != nil {
			fmt.Printf("failed to fetch address for hotelID=%v with err=%v", hotelID, err)
		} else {
			location.Address = address
		}

		city, err := getCity(hotelID)
		if err != nil {
			fmt.Printf("failed to fetch city for hotelID=%v with err=%v", hotelID, err)
		} else {
			location.City = city
		}

		country, err := getCountry(hotelID)
		if err != nil {
			fmt.Printf("failed to fetch country for hotelID=%v with err=%v", hotelID, err)
		} else {
			location.Country = country
		}

		hotel.Location = location
	}
}

func PrepareDataSource() {
	buildDataSource()
	buildDataSourceOrderingForField()
	buildDataset()
}
