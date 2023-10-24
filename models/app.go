package model

import (
	"pm/domain"

	"github.com/golang-jwt/jwt/v5"
)

var Header interface{}

type ServerConfig struct {
	Name      string `env:"APP_NAME"`
	Port      string `env:"APP_PORT"`
	Host      string `env:"APP_HOST"`
	JWTSecret string `ennv:"JWT_SECRET"`
	DBConfig  DBConfig
}

// db primary config
type DBConfig struct {
	Name     string `env:"DB_NAME"`
	Host     string `env:"DB_HOST"`
	Port     string `env:"DB_PORT"`
	User     string `env:"DB_USER"`
	Password string `env:"DB_PASS"`
}

type Claims struct {
	jwt.RegisteredClaims
	User domain.User `json:"user"`
}
