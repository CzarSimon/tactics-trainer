package config

import (
	"fmt"
	"strings"
	"time"

	"github.com/CzarSimon/httputil/dbutil"
	"github.com/CzarSimon/httputil/environ"
	"github.com/CzarSimon/httputil/jwt"
	"github.com/CzarSimon/httputil/logger"
)

var log = logger.GetDefaultLogger("internal/config").Sugar()

// By default a token is configured to be valid for a week.
const defaultTokenLifetime = "168h"

// Config application configuration.
type Config struct {
	DB             dbutil.Config
	MigrationsPath string
	Port           string
	JwtCredentials jwt.Credentials
	TokenLifetime  time.Duration
	KEKPath        string
}

// GetConfig reads, parses and marshalls the applications configuration.
func GetConfig() Config {
	return Config{
		DB:             getDBConfig(),
		MigrationsPath: environ.Get("MIGRATIONS_PATH", "/etc/iam-server/db/sqlite"),
		Port:           environ.Get("PORT", "8080"),
		JwtCredentials: jwt.Credentials{
			Issuer: environ.Get("JWT_ISSUER", "tactics-trainer/iam-server"),
			Secret: environ.MustGet("JWT_SECRET"),
		},
		TokenLifetime: getTokenLifetime(),
		KEKPath:       environ.Get("KEY_ENCRYPTION_KEYS_PATH", "/etc/iam-server/key-encryption-keys.txt"),
	}
}

func getDBConfig() dbutil.Config {
	dbType := strings.ToLower(environ.Get("DB_TYPE", "mysql"))
	if dbType == "sqlite" {
		return dbutil.SqliteConfig{
			Name: environ.MustGet("DB_FILENAME"),
		}
	}

	tlsConfig := fmt.Sprintf("tls=%s", environ.Get("DB_SSL_MODE", "true"))
	return dbutil.MysqlConfig{
		Host:             environ.MustGet("DB_HOST"),
		Port:             environ.Get("DB_PORT", "3306"),
		Database:         environ.Get("DB_NAME", "iam-server"),
		User:             environ.MustGet("DB_USERNAME"),
		Password:         environ.MustGet("DB_PASSWORD"),
		ConnectionParams: fmt.Sprintf("parseTime=true&%s", tlsConfig),
	}
}

func getTokenLifetime() time.Duration {
	lifetimeStr := environ.Get("TOKEN_LIFETIME", defaultTokenLifetime)
	lifetime, err := time.ParseDuration(lifetimeStr)
	if err != nil {
		log.Panic("Failed to parse token lifetime %s. Error %w", lifetimeStr, err)
	}

	return lifetime
}
