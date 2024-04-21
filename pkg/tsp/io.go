package tsp

import (
	"bufio"
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func checkErr(e error) {
	if e != nil {
		panic(e)
	}
}

func ReadCsv(opts *Opts) *[][]string {
	var file io.Reader
	var err error
	if opts.File == "" {
		file = bufio.NewReader(os.Stdin)
	} else {
		fp, err := filepath.Abs(opts.File)
		checkErr(err)

		file, err = os.Open(fp)
		checkErr(err)
	}

	csvReader := csv.NewReader(file)
	data, err := csvReader.ReadAll()
	checkErr(err)

	return &data
}

func ParseDDFromRecord(record []string) (DDCoordinates, error) {
	var latDegrees, longDegrees float64

	if len(record) != 3 {
		err := errors.New("every row must containt location name, latitude, and longitude")
		return DDCoordinates{}, err
	}

	latDegrees, err := ParseFloat(record[1])
	if err != nil {
		return DDCoordinates{}, err
	}

	longDegrees, err = ParseFloat(record[2])
	if err != nil {
		return DDCoordinates{}, err
	}

	return DDCoordinates{record[0], latDegrees, longDegrees}, nil
}

func RecordsToCoordinates(records *[][]string) *[]DDCoordinates {
	coordinates := make([]DDCoordinates, len(*records))
	var err error
	for i, record := range *records {
		coordinates[i], err = ParseDDFromRecord(record)
		checkErr(err)
	}
	return &coordinates
}

func ParseFloat(s string) (float64, error) {
	var floatRes float64
	_, err := fmt.Sscanf(s, "%f", &floatRes)
	if err != nil {
		return 0, err
	}

	return floatRes, nil
}
