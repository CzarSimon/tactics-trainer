package config

import (
	"fmt"
	"strings"

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
}

// GetConfig reads, parses and marshalls the applications configuration.
func GetConfig() Config {
	return Config{
		DB:             getDBConfig(),
		MigrationsPath: environ.Get("MIGRATIONS_PATH", "/etc/puzzle-service/db/mysql"),
		Port:           environ.Get("PORT", "8080"),
		JwtCredentials: jwt.Credentials{
			Issuer: environ.Get("JWT_ISSUER", "tactics-trainer/iam-server"),
			Secret: environ.MustGet("JWT_SECRET"),
		},
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
		Database:         environ.Get("DB_NAME", "puzzle-server"),
		User:             environ.MustGet("DB_USERNAME"),
		Password:         environ.MustGet("DB_PASSWORD"),
		ConnectionParams: fmt.Sprintf("parseTime=true&%s", tlsConfig),
	}
}
