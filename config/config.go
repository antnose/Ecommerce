package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var configurations Config

type Config struct {
	Version     string
	ServiceName string
	HttpPort    int
}

func loadConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Failed to load env", err)
		os.Exit(1)
	}

	version := os.Getenv("VERSION")
	if version == "" {
		log.Fatal("Version is required")
		os.Exit(1)
	}

	serviceName := os.Getenv("SERVICE_NAME")
	if serviceName == "" {
		log.Fatal("Service Name is required")
		os.Exit(1)
	}

	httpPort := os.Getenv("HTTP_PORT")

	if httpPort == "" {
		log.Fatal("HTTP Port is required")
		os.Exit(1)
	}

	port, err := strconv.ParseInt(httpPort, 10, 64)
	if err != nil {
		log.Fatal("Provide a valid port")
		os.Exit(1)
	}

	configurations = Config{
		Version:     version,
		ServiceName: serviceName,
		HttpPort:    int(port),
	}
}

func GetConfig() Config {
	loadConfig()
	return configurations
}
