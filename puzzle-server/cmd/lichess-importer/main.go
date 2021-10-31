package main

import (
	"context"
	"io"

	"github.com/CzarSimon/httputil/logger"
	_ "github.com/go-sql-driver/mysql"
)

var log = logger.GetDefaultLogger("cmd/importer").Sugar()

func main() {
	repo, db := setupPuzzleRepository()
	defer db.Close()

	filename := input("Puzzle file")
	reader := NewPuzzleReader(filename)
	defer reader.Close()

	ctx := context.Background()
	for {
		puzzle, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		err = repo.Save(ctx, puzzle)
		if err != nil {
			log.Fatal(err)
		}
	}
}
