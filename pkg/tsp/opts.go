package tsp

import (
	"github.com/hellflame/argparse"
)

type Opts struct {
	File                 string
	UseEuclideanDistance bool
}

func GetOpts() (*Opts, error) {
	parser := argparse.NewParser("tsp", "get all command line values", &argparse.ParserConfig{
		DisableDefaultShowHelp: true,
	})

	file := parser.String("f", "file", &argparse.Option{
		Positional: true,
		Required:   false,
		Default:    "",
	})

	euclidean := parser.Flag("e", "euclidean", &argparse.Option{
		Positional: false,
		Required:   false,
		Default:    "",
	})

	err := parser.Parse(nil)
	if err != nil {
		return nil, err
	}

	return &Opts{
		File:                 *file,
		UseEuclideanDistance: *euclidean,
	}, nil
}
