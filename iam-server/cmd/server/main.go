package main

import (
	"github.com/CzarSimon/tactics-trainer/iam-server/internal/api"
	"github.com/CzarSimon/tactics-trainer/iam-server/internal/config"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	cfg := config.GetConfig()
	api.Start(cfg)
}
