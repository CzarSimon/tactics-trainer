package main

import (
	"encoding/csv"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/CzarSimon/httputil/id"
	"github.com/CzarSimon/httputil/timeutil"
	"github.com/CzarSimon/tactics-trainer/puzzle-server/internal/models"
)

type PuzzleReader struct {
	f   io.ReadCloser
	csv *csv.Reader
}

func NewPuzzleReader(filename string) PuzzleReader {
	f, err := os.Open(filepath.Clean(filename))
	if err != nil {
		log.Fatal(err)
	}

	csvReader := csv.NewReader(f)
	_, err = csvReader.Read()
	if err != nil {
		log.Fatal(err)
	}

	return PuzzleReader{
		f:   f,
		csv: csvReader,
	}
}

func (r PuzzleReader) Read() (models.Puzzle, error) {
	row, err := r.csv.Read()
	if err != nil {
		return models.Puzzle{}, err
	}

	return models.Puzzle{
		ID:              id.New(),
		ExternalID:      row[0],
		FEN:             row[1],
		Moves:           strings.Split(row[2], " "),
		Rating:          parseInt(row[3]),
		RatingDeviation: parseInt(row[4]),
		Popularity:      parseInt(row[5]),
		Themes:          strings.Split(row[7], " "),
		GameURL:         row[8],
		CreatedAt:       timeutil.Now(),
		UpdatedAt:       timeutil.Now(),
	}, nil
}

func (r PuzzleReader) Close() error {
	return r.f.Close()
}

func parseInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}

	return i
}
