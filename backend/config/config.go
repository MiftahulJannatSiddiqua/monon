package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var configuration *Config

type DBConfig struct {
	User          string
	Password      string
	Host          string
	Port          int
	Name          string
	EnableSslmode bool
}

type Config struct {
	Version      string
	ServiceName  string
	HttpPort     int
	JwtSecretKey string
	DB           *DBConfig
}

func loadConfig() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("faild to load env", err)
		os.Exit(1)
	}

	version := os.Getenv("VERSION")
	if version == "" {
		fmt.Println("version is required")
		os.Exit(1)
	}

	serviceName := os.Getenv("SERVICE_NAME")
	if serviceName == "" {
		fmt.Println("serviceName is required")
		os.Exit(1)
	}

	httpPort := os.Getenv("HTTP_PORT")
	if httpPort == "" {
		fmt.Println("httpPort is required")
		os.Exit(1)
	}

	port, err := strconv.ParseInt(httpPort, 10, 64)
	if err != nil {
		fmt.Println("port must be number")
		os.Exit(1)
	}

	jwtSecretKey := os.Getenv("JWT_SECRET_KEY")
	if jwtSecretKey == "" {
		fmt.Println("secretKey is required")
		os.Exit(1)
	}

	dbUser := os.Getenv("DB_USER")
	if dbUser == "" {
		fmt.Println("dbUser is required")
		os.Exit(1)
	}

	dbPassword := os.Getenv("DB_PASSWORD")
	if dbPassword == "" {
		fmt.Println("dbPassword is required")
		os.Exit(1)
	}

	dbHost := os.Getenv("DB_HOST")
	if dbHost == "" {
		fmt.Println("dbHost is required")
		os.Exit(1)
	}

	dbPort := os.Getenv("DB_PORT")
	if dbPort == "" {
		fmt.Println("dbPort is required")
		os.Exit(1)
	}

	dbPrt, err := strconv.ParseInt(dbPort, 10, 64)
	if err != nil {
		fmt.Println("port must be number")
		os.Exit(1)
	}

	dbName := os.Getenv("DB_NAME")
	if dbName == "" {
		fmt.Println("dbName is required")
		os.Exit(1)
	}

	dbEnableSslMode := os.Getenv("DB_ENABLE_SSL_MODE")
	if dbEnableSslMode == "" {
		fmt.Println("dbEnableSslMode is required")
		os.Exit(1)
	}

	dbEnableSSLMode, err := strconv.ParseBool(dbEnableSslMode)
	if err != nil {
		fmt.Println("invlaid dbEnableSSLMode vale", err)
		os.Exit(1)
	}

	DbConfig := &DBConfig{
		User:          dbUser,
		Password:      dbPassword,
		Host:          dbHost,
		Port:          int(dbPrt),
		Name:          dbName,
		EnableSslmode: dbEnableSSLMode,
	}

	configuration = &Config{
		Version:      version,
		ServiceName:  serviceName,
		HttpPort:     int(port), // type casting
		JwtSecretKey: jwtSecretKey,
		DB:           DbConfig,
	}
}

func GetConfig() *Config {
	if configuration == nil {
		//first time
		loadConfig()
	}
	return configuration
}

