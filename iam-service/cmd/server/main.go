package main

import (
	"github.com/CzarSimon/tactics-trainer/iam-service/internal/api"
	"github.com/CzarSimon/tactics-trainer/iam-service/internal/config"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	cfg := config.GetConfig()
	api.Start(cfg)
}
