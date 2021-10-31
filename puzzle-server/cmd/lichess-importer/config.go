package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"os"
	"strings"

	"github.com/CzarSimon/httputil/dbutil"
	"github.com/CzarSimon/tactics-trainer/puzzle-server/internal/config"
	"github.com/CzarSimon/tactics-trainer/puzzle-server/internal/repository"
)

func setupPuzzleRepository() (repository.PuzzleRepository, *sql.DB) {
	cfg := config.GetConfig()
	db := dbutil.MustConnect(cfg.DB)
	err := dbutil.Upgrade(cfg.MigrationsPath, cfg.DB.Driver(), db)
	if err != nil {
		log.Fatal(err)
	}

	return repository.NewPuzzleRepository(db), db
}

func input(prompt string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("%s: ", prompt)
	text, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	return strings.ReplaceAll(text, "\n", "")
}
