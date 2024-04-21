package tsp

import (
	"math"
)

type DDCoordinates struct {
	Name      string
	Latitude  float64
	Longitude float64
}

func PathDist(path []DDCoordinates, distFunc func(p1, p2 DDCoordinates) float64) float64 {
	distance := 0.0
	n := len(path)
	for i := 0; i < n-1; i++ {
		distance += distFunc(path[i], path[i+1])
	}
	distance += distFunc(path[n-1], path[0])
	return distance
}

const earthRadiusMiles = 3958.8 // Earth radius in miles

func Haversine(p1, p2 DDCoordinates) float64 {
	// calculates the great circlue distance between two points on a sphere

	// Convert latitude and longitude from degrees to radians
	lat1 := degToRad(p1.Latitude)
	lon1 := degToRad(p1.Longitude)
	lat2 := degToRad(p2.Latitude)
	lon2 := degToRad(p2.Longitude)

	latDiff := lat2 - lat1
	lonDiff := lon2 - lon1

	// Haversine formula
	a := math.Pow(math.Sin(latDiff/2), 2) + math.Cos(lat1)*math.Cos(lat2)*math.Pow(math.Sin(lonDiff/2), 2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	distance := earthRadiusMiles * c

	return distance
}

func degToRad(deg float64) float64 {
	return deg * (math.Pi / 180)
}

func EuclideanDistance(p1 DDCoordinates, p2 DDCoordinates) float64 {
	xDiff := math.Abs(p1.Latitude - p2.Latitude)
	yDiff := math.Abs(p1.Longitude - p2.Longitude)
	return math.Sqrt(math.Pow(xDiff, 2) + math.Pow(yDiff, 2))
}
