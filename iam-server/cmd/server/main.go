package main

import (
	"github.com/CzarSimon/tactics-trainer/iam-server/internal/api"
	"github.com/CzarSimon/tactics-trainer/iam-server/internal/config"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	cfg := config.GetConfig()
	api.Start(cfg)
}
