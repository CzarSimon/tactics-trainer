package api

import (
	"database/sql"
	"net/http"

	"github.com/CzarSimon/httputil"
	"github.com/CzarSimon/httputil/dbutil"
	"github.com/CzarSimon/httputil/logger"
	"github.com/CzarSimon/tactics-trainer/iam-service/internal/config"
	"go.uber.org/zap"
)

var log = logger.GetDefaultLogger("internal/api")

func Start(cfg config.Config) {
	db := dbutil.MustConnect(cfg.DB)
	defer db.Close()

	err := dbutil.Upgrade(cfg.MigrationsPath, cfg.DB.Driver(), db)
	if err != nil {
		log.Panic("Failed to apply upgrade migratons", zap.Error(err))
	}

	r := httputil.NewRouter("iam-service", healthCheck(db))

	server := &http.Server{
		Addr:    ":" + cfg.Port,
		Handler: r,
	}

	log.Info("Starting iam-service", zap.String("port", cfg.Port))
	err = server.ListenAndServe()
	if err != nil {
		log.Error("Server stoped with an error", zap.Error(err))
	}
}

func healthCheck(db *sql.DB) httputil.HealthFunc {
	return func() error {
		return dbutil.Connected(db)
	}
}
