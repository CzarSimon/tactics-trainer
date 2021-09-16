package main

import (
	"context"

	"github.com/CzarSimon/httputil/dbutil"
	"github.com/CzarSimon/httputil/environ"
	"github.com/CzarSimon/httputil/logger"
	"github.com/CzarSimon/httputil/timeutil"
	"github.com/CzarSimon/tactics-trainer/puzzle-server/internal/models"
	"github.com/CzarSimon/tactics-trainer/puzzle-server/internal/repository"
	_ "github.com/mattn/go-sqlite3"
)

var log = logger.GetDefaultLogger("cmd/importer").Sugar()

func main() {
	puzzle := models.Puzzle{
		ID:         "puzzle-0",
		ExternalID: "ext-id-0",
		FEN:        "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1",
		Moves: []string{
			"e2e4",
			"e7e5",
			"g1f3",
		},
		Rating:          1500,
		RatingDeviation: 100,
		Popularity:      10,
		Themes: []string{
			"mateIn2",
			"short",
			"sacrifice",
		},
		GameURL:   "https://some.url/some/id",
		CreatedAt: timeutil.Now(),
		UpdatedAt: timeutil.Now(),
	}

	repo := setupPuzzleRepository()
	ctx := context.Background()
	err := repo.Save(ctx, puzzle)
	if err != nil {
		log.Fatal(err)
	}
}

func setupPuzzleRepository() repository.PuzzleRepository {
	cfg := getDBConfig()
	db := dbutil.MustConnect(cfg)
	migrationsPath := environ.MustGet("MIGRATIONS_PATH")
	dbutil.Upgrade(migrationsPath, cfg.Driver(), db)
	return repository.NewPuzzleRepository(db)
}

func getDBConfig() dbutil.Config {
	return dbutil.SqliteConfig{
		Name: environ.MustGet("DB_NAME"),
	}
}
