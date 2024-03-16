package main

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

var (
	API_ENDPOINT = "https://nominatim.openstreetmap.org/search.php?polygon_geojson=1&format=jsonv2&q="
)

func getGeoCodingAPI(endpoint string) (float64, float64, error) {
	response, err := http.Get(endpoint)
	if err != nil {
		return 0.0, 0.0, err
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return 0.0, 0.0, err
	}

	var data []interface{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		return 0.0, 0.0, err
	}

	averageLat := 0.0
	averageLon := 0.0
	for _, v := range data {
		// calculate average lat and lon because in vietnam, we not using zipcode for each place
		// so we need to calculate average lat and lon to get the center of the place
		averageLat += v.(map[string]interface{})["lat"].(float64)
		averageLon += v.(map[string]interface{})["lon"].(float64)
	}

	averageLat /= float64(len(data))
	averageLon /= float64(len(data))

	return averageLat, averageLon, nil
}

func GeoCodingByAddress(address string) (result Coordinate, err error) {
	endpoint := API_ENDPOINT + address
	lat, lon, _ := getGeoCodingAPI(endpoint)
	if lat == 0 && lon == 0 {
		err = errors.New("cannot find the address")
	}

	return Coordinate{Lat: lat, Lon: lon}, err
}

func GeoCodingByZipcode(zipcode string) (result Coordinate, err error) {
	endpoint := API_ENDPOINT + zipcode
	lat, lon, err := getGeoCodingAPI(endpoint)
	if lat == 0 && lon == 0 {
		err = errors.New("cannot find the zipcode")
	}

	return Coordinate{Lat: lat, Lon: lon}, err
}
