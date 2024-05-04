package patagonia

import (
	"ascenda/common"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
)

type images struct {
	URL         string `json:"url"`
	Description string `json:"description"`
}

type hotel struct {
	ID          string               `json:"id"`
	Destination int64                `json:"destination"`
	Name        string               `json:"name"`
	Latitude    float64              `json:"lat"`
	Longitude   float64              `json:"lng"`
	Address     string               `json:"address"`
	Information string               `json:"info"`
	Amenities   []string             `json:"amenities"`
	Images      map[string][]*images `json:"images"`
}

const (
	supplierURL = "https://5f2be0b4ffc88500167b85a0.mockapi.io/suppliers/patagonia"
)

var (
	Source common.Supplier
	once   sync.Once
)

type DataSource struct {
	dataSource map[string]*hotel
	hotelIDs   []string
}

func fetchData() ([]*hotel, error) {
	response, err := http.Get(supplierURL)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var resp []*hotel
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func prepareDataSource(hotels []*hotel) *DataSource {
	source := DataSource{
		dataSource: map[string]*hotel{},
		hotelIDs:   []string{},
	}
	for _, hotelData := range hotels {
		source.dataSource[hotelData.ID] = hotelData
		source.hotelIDs = append(source.hotelIDs, hotelData.ID)
	}
	return &source
}

func Init() (common.Supplier, error) {
	once.Do(func() {
		resp, err := fetchData()
		if err != nil {
			fmt.Printf("failed to fetch data with err: %v", err)
		}
		datasource := prepareDataSource(resp)
		Source = datasource
	})
	if Source == nil {
		return nil, fmt.Errorf("failed to fetch data from acme")
	}
	return Source, nil
}
