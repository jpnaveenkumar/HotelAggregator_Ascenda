package handler

import (
	"ascenda/common"
	"ascenda/logic"
)

func filterHotelsByDestinationID(hotels []*common.Hotel, destinationID int64) []*common.Hotel {
	var response []*common.Hotel
	for _, hotel := range hotels {
		if hotel.DestinationID == destinationID {
			response = append(response, hotel)
		}
	}
	return response
}

func fetchHotels(hotelIDs []string, destinationID int64) []*common.Hotel {
	var hotels []*common.Hotel
	if len(hotelIDs) > 0 {
		hotels = logic.GetHotels(hotelIDs)
		if destinationID != 0 {
			hotels = filterHotelsByDestinationID(hotels, destinationID)
		}
		return hotels
	}
	return logic.GetHotelsByDestinationID(destinationID)
}
