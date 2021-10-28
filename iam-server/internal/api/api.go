package api

import (
	"database/sql"
	"net/http"

	"github.com/CzarSimon/httputil"
	"github.com/CzarSimon/httputil/crypto"
	"github.com/CzarSimon/httputil/dbutil"
	"github.com/CzarSimon/httputil/jwt"
	"github.com/CzarSimon/httputil/logger"
	"github.com/CzarSimon/tactics-trainer/iam-server/internal/api/authentication"
	"github.com/CzarSimon/tactics-trainer/iam-server/internal/config"
	"github.com/CzarSimon/tactics-trainer/iam-server/internal/repository"
	"github.com/CzarSimon/tactics-trainer/iam-server/internal/service"
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

	cipher := &service.Cipher{
		KEKRepo: repository.NewKeyEncryptionKeyRepository(config.LoadKEKConfig(cfg.KEKPath), db),
	}

	authSvc := &service.AuthenticationService{
		UserRepo:      repository.NewUserRepository(db),
		Cipher:        cipher,
		Hasher:        crypto.DefaultScryptHasher(),
		Issuer:        jwt.NewIssuer(cfg.JwtCredentials),
		TokenLifetime: cfg.TokenLifetime,
	}

	r := httputil.NewRouter("iam-server", healthCheck(db))
	authentication.AttachController(authSvc, r)

	server := &http.Server{
		Addr:    ":" + cfg.Port,
		Handler: r,
	}

	log.Info("Starting iam-server", zap.String("port", cfg.Port))
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
