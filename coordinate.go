package main

import (
	"fmt"
	"math"
)

type Coordinate struct {
	Lat float64 `json:"lat"`
	Lon float64 `json:"lon"`
}

func (c Coordinate) GetLat() float64 {
	return c.Lat
}

func (c Coordinate) GetLong() float64 {
	return c.Lon
}

func (c *Coordinate) SetLat(lat float64) {
	c.Lat = lat
}

func (c *Coordinate) SetLong(lon float64) {
	c.Lon = lon
}

func (c Coordinate) Valid() bool {
	return c.Lat != 0 && c.Lon != 0
}

func (c Coordinate) Empty() bool {
	return c.Lat == 0 && c.Lon == 0
}

func (c Coordinate) Equal(other Coordinate) bool {
	return c.Lat == other.Lat && c.Lon == other.Lon
}

func (c Coordinate) String() string {
	return fmt.Sprintf("%f,%f", c.Lat, c.Lon)
}

// using Euclidean distance
func (c Coordinate) CalculateEuclideanDistance(inputLat float64, inputLon float64) float64 {
	return math.Sqrt((c.Lat-inputLat)*(c.Lat-inputLat) + (c.Lon-inputLon)*(c.Lon-inputLon))
}

// using distance between two points on the surface of earth
func (c Coordinate) CalculateHaversineDistance(inputLat float64, inputLon float64) float64 {
	// convert decimal degrees to radians
	lat1 := c.Lat * math.Pi / 180
	lat2 := inputLat * math.Pi / 180
	lon1 := c.Lon * math.Pi / 180
	lon2 := inputLon * math.Pi / 180

	// haversine formula
	dlon := lon2 - lon1
	dlat := lat2 - lat1
	a := math.Pow(math.Sin(dlat/2), 2) + math.Cos(lat1)*math.Cos(lat2)*math.Pow(math.Sin(dlon/2), 2)
	cen := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	r := 6371.0 // Radius of earth in kilometers
	return r * cen
}

// calculate the distance between two points on the surface of earth, then getting the distance within a threshold
func (c Coordinate) GetNearbyCoordinates(threshold int, usingEuclidean bool, inputCoordinate ...Coordinate) []Coordinate {
	nearbyCoordinates := make([]Coordinate, 0)
	distance := 0.0
	for _, v := range inputCoordinate {
		if usingEuclidean {
			distance = c.CalculateEuclideanDistance(v.Lat, v.Lon)
		} else {
			distance = c.CalculateHaversineDistance(v.Lat, v.Lon)

		}
		if distance <= float64(threshold) {
			nearbyCoordinates = append(nearbyCoordinates, v)
		}
	}

	return nearbyCoordinates
}
