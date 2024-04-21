package main

import (
	"fmt"
	"log"

	"github.com/regular-polygon/goTravel/pkg/tsp"
)

func main() {
	opts, err := tsp.GetOpts()
	if err != nil {
		log.Fatal("Unable to get options\n")
	}

	data := tsp.ReadCsv(opts)
	coordinates := tsp.RecordsToCoordinates(data)

	var bestPath []tsp.DDCoordinates
	var bestDist float64
	if opts.UseEuclideanDistance {
		bestPath, bestDist = tsp.BruteForceTSP(*coordinates, tsp.EuclideanDistance)
	} else {
		bestPath, bestDist = tsp.BruteForceTSP(*coordinates, tsp.Haversine)
	}

	for _, p := range bestPath {
		fmt.Println(p)
	}
	fmt.Println(bestDist)
}
