package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"os"
	"strings"

	"github.com/CzarSimon/httputil/dbutil"
	"github.com/CzarSimon/httputil/environ"
	"github.com/CzarSimon/tactics-trainer/puzzle-server/internal/repository"
)

func setupPuzzleRepository() (repository.PuzzleRepository, *sql.DB) {
	cfg := getDBConfig()
	db := dbutil.MustConnect(cfg)
	migrationsPath := input("migrations path")
	err := dbutil.Upgrade(migrationsPath, cfg.Driver(), db)
	if err != nil {
		log.Fatal(err)
	}

	return repository.NewPuzzleRepository(db), db
}

func getDBConfig() dbutil.Config {
	dbType := environ.Get("DB_TYPE", "mysql")
	if dbType == "sqlite" {
		return dbutil.SqliteConfig{
			Name: input("Database name"),
		}
	}

	return dbutil.MysqlConfig{
		Host:     input("host"),
		Port:     input("port"),
		User:     input("user"),
		Password: input("password"),
		Database: input("database"),
	}
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
