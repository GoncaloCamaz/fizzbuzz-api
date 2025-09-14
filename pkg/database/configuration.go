package database

import (
	"net/url"
	"runtime"
	"strings"
)

// Config represents the configuration for a database
type Config struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
	PoolSize int
}

// NewDBConfig creates a new database configuration with custom values
func NewDBConfig(host, port, user, password, database string) *Config {
	return &Config{
		Host:     host,
		Port:     port,
		User:     user,
		Password: password,
		Database: database,
		PoolSize: 4 * runtime.GOMAXPROCS(0),
	}
}

// GetDSN returns the connection DSN string
func (c Config) GetDSN() string {
	builder := strings.Builder{}

	_, _ = builder.WriteString("postgres://")
	_, _ = builder.WriteString(url.QueryEscape(c.User))
	_, _ = builder.WriteRune(':')
	_, _ = builder.WriteString(url.QueryEscape(c.Password))
	_, _ = builder.WriteRune('@')
	_, _ = builder.WriteString(c.Host)
	_, _ = builder.WriteRune(':')
	_, _ = builder.WriteString(c.Port)
	_, _ = builder.WriteRune('/')
	_, _ = builder.WriteString(c.Database)

	_, _ = builder.WriteString("?sslmode=disable")

	return builder.String()
}
