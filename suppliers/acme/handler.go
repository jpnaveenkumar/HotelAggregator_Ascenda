package acme

import (
	"ascenda/common"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
)

const (
	supplierURL = "https://5f2be0b4ffc88500167b85a0.mockapi.io/suppliers/acme"
)

var (
	Source common.Supplier
	once   sync.Once
)

type DataSource struct {
	dataSource map[string]*hotel
	hotelIDs   []string
}

type hotel struct {
	Id            string
	DestinationId int64
	Name          string
	Latitude      interface{}
	Longitude     interface{}
	Address       string
	City          string
	Country       string
	PostalCode    string
	Description   string
	Facilities    []string
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
		source.dataSource[hotelData.Id] = hotelData
		source.hotelIDs = append(source.hotelIDs, hotelData.Id)
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
