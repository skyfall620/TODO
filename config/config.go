package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	App  AppConfig
	Db   DbConfig
	Auth AuthConfig
}

type AuthConfig struct {
	Secret string
}

type AppConfig struct {
	HTTPAddr      string
	ReadTimeout   int
	WriteTimeout  int
	MaxHeaderByte int
}

type DbConfig struct {
	DBHost     string
	DBPort     int
	DBUser     string
	DBPassword string
	DBName     string
	SSLMode    string
}

func MustLoadConfig() *Config {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Warning: .env file not loaded: %s", err.Error())
	}

	cfg := &Config{}

	cfg.Auth = AuthConfig{
		Secret: os.Getenv("JWT_SECRET"),
	}

	cfg.App = AppConfig{
		HTTPAddr: os.Getenv("HTTP_ADDR"),
	}

	readTimeout, err := strconv.Atoi(os.Getenv("READ_TIMEOUT"))
	if err != nil {
		log.Printf("Warning: READ_TIMEOUT is not a valid integer, using default 10: %v", err)
		readTimeout = 10
	}
	cfg.App.ReadTimeout = readTimeout

	writeTimeout, err := strconv.Atoi(os.Getenv("WRITE_TIMEOUT"))
	if err != nil {
		log.Printf("Warning: WRITE_TIMEOUT is not a valid integer, using default 10: %v", err)
		writeTimeout = 10
	}
	cfg.App.WriteTimeout = writeTimeout

	maxHeaderByte, err := strconv.Atoi(os.Getenv("MAX_HEADER_BYTE"))
	if err != nil {
		log.Printf("Warning: MAX_HEADER_BYTE is not a valid integer, using default 20: %v", err)
		maxHeaderByte = 20
	}
	cfg.App.MaxHeaderByte = maxHeaderByte

	dbPort, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		log.Printf("Warning: DB_PORT is not a valid integer, using default 5432: %v", err)
		dbPort = 5432
	}

	cfg.Db = DbConfig{
		DBHost:     os.Getenv("DB_HOST"),
		DBPort:     dbPort,
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBName:     os.Getenv("DB_NAME"),
		SSLMode:    os.Getenv("DB_SSLMODE"),
	}

	return cfg
}

func (c Config) DBConnString() string {

	return fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=%s",
		c.Db.DBUser, c.Db.DBPassword, c.Db.DBHost, c.Db.DBPort, c.Db.DBName, c.Db.SSLMode,
	)
}
