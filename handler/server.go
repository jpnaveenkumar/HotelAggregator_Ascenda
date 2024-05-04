package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

func getErrorResponse(err string) []byte {
	errResp := map[string]map[string]string{
		"data": {
			"error": err,
		},
	}

	byteArr, _ := json.Marshal(errResp)
	return byteArr
}

func getSuccessResponse(data interface{}) []byte {
	resp := map[string]interface{}{
		"data": data,
	}

	byteArr, _ := json.Marshal(resp)
	return byteArr
}

func fetchHotelsHandler(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	queryParams := request.URL.Query()

	hotelIDString := queryParams.Get("hotelIDs")
	destinationID := queryParams.Get("destinationIDs")

	if len(hotelIDString) == 0 && len(destinationID) == 0 {
		writer.WriteHeader(http.StatusBadRequest)
		writer.Write(getErrorResponse("hotelIDs or destinationID is mandatory"))
		return
	}

	var destID int64
	if len(destinationID) > 0 {
		id, err := strconv.Atoi(destinationID)
		if err != nil {
			writer.WriteHeader(http.StatusBadRequest)
			writer.Write(getErrorResponse("invalid destinationID"))
			return
		}
		destID = int64(id)
	}

	hotelIDs := strings.Split(hotelIDString, ",")
	var filteredHotelIDs []string
	for _, hotelID := range hotelIDs {
		if len(hotelID) > 0 {
			filteredHotelIDs = append(filteredHotelIDs, strings.TrimSpace(hotelID))
		}
	}

	hotels := fetchHotels(filteredHotelIDs, destID)
	if len(hotels) == 0 {
		writer.WriteHeader(http.StatusNoContent)
		return
	} else {
		writer.WriteHeader(http.StatusOK)
		writer.Write(getSuccessResponse(hotels))
	}
}

func StartHTTPServer() {

	basePath := "/api/v1"

	http.HandleFunc(basePath+"/hotels", fetchHotelsHandler)

	fmt.Println("Starting HTTP server at port 8080...")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("error starting server")
	}
}
