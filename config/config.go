package config

import (
	"fmt"
	"time"
)

type (
	Database struct {
		Host           string        `mapstructure:"host" validate:"required"`
		Port           int           `mapstructure:"port" validate:"required"`
		User           string        `mapstructure:"user" validate:"required"`
		Password       string        `mapstructure:"password" validate:"required"`
		DBName         string        `mapstructure:"dbname" validate:"required"`
		SSLMode        string        `mapstructure:"sslmode" validate:"required"`
		Schema         string        `mapstructure:"schema" validate:"required"`
		ConnectTimeout time.Duration `mapstructure:"connectTimeout" validate:"required"`
		ReadTimeout    time.Duration `mapstructure:"readTimeout" validate:"required"`
	}

	Server struct {
		Port         int           `mapstructure:"port" validate:"required"`
		Timeout      time.Duration `mapstructure:"timeout" validate:"required"`
		BodyLimit    string        `mapstructure:"bodyLimit" validate:"required"`
		AllowOrigins []string      `mapstructure:"allowOrigins" validate:"required"`
		LogLevel     string        `mapstructure:"logLevel" validate:"required"`
	}

	Config struct {
		Database *Database `mapstructure:"database" validate:"required"`
		Server   *Server   `mapstructure:"server" validate:"required"`
		Mailer   *Mailer   `mapstructure:"mailer" validate:"required"`
	}

	Mailer struct {
		Host     string `mapstructure:"host" validate:"required"`
		Port     int    `mapstructure:"port" validate:"required"`
		Username string `mapstructure:"username" validate:"required"`
		Password string `mapstructure:"password" validate:"required"`
		From     string `mapstructure:"from" validate:"required"`
	}
)

func (d *Database) GetDSN() string {
	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=%s search_path=%s",
		d.Host,
		d.User,
		d.Password,
		d.DBName,
		d.Port,
		d.SSLMode,
		d.Schema,
	)
}

func (d *Database) GetConnectionString() string {
	return fmt.Sprintf("postgresql://%s:%s@%s:%d/%s?sslmode=%s", d.User, d.Password, d.Host, d.Port, d.DBName, d.SSLMode)
}
