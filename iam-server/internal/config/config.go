package config

import (
	"strings"
	"time"

	"github.com/CzarSimon/httputil/dbutil"
	"github.com/CzarSimon/httputil/environ"
	"github.com/CzarSimon/httputil/jwt"
)

// Config application configuration.
type Config struct {
	DB             dbutil.Config
	MigrationsPath string
	Port           string
	JwtCredentials jwt.Credentials
	TokenLifetime  time.Duration
}

// GetConfig reads, parses and marshalls the applications configuration.
func GetConfig() Config {
	return Config{
		DB:             getDBConfig(),
		MigrationsPath: environ.Get("MIGRATIONS_PATH", "/etc/iam-server/db/sqlite"),
		Port:           environ.Get("PORT", "8080"),
	}
}

func getDBConfig() dbutil.Config {
	dbType := strings.ToLower(environ.Get("DB_TYPE", "mysql"))
	if dbType == "sqlite" {
		return dbutil.SqliteConfig{
			Name: environ.MustGet("DB_FILENAME"),
		}
	}

	return dbutil.MysqlConfig{
		Host:             environ.MustGet("DB_HOST"),
		Port:             environ.Get("DB_PORT", "3306"),
		Database:         environ.Get("DB_NAME", "iam-server"),
		User:             environ.MustGet("DB_USERNAME"),
		Password:         environ.MustGet("DB_PASSWORD"),
		ConnectionParams: "parseTime=true",
	}
}
