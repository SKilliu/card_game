package server

import (
	"os"
	"strconv"
)

type Configuration struct {
	Host string `env:"host"`
	Port int    `env:"port"`
}

var configuration Configuration

func loadConfigFromEnvs() (err error) {
	port, ok := os.LookupEnv("port")
	if !ok {
		port = "8080"
	}
	configuration.Port, err = strconv.Atoi(port)
	if err != nil {
		return
	}

	host, ok := os.LookupEnv("host")
	if !ok {
		host = "localhost"
	}
	configuration.Host = host

	return
}
