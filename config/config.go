package config

import (
	"os"

	"go.uber.org/zap"
)

type Configuration struct {
	service  Service
	database Database
}

type Database struct {
	Host     string
	Port     string
	Username string
	Password string
	Schema   string
}

type Service struct {
	Name        string
	Port        string
	Servicehost string
}

func NewConfig(log *zap.Logger) *Configuration {
	var config Configuration
	config.database.Host = os.Getenv("DBHOST")
	if len(config.database.Host) == 0 {
		log.Fatal("No value in env DBHOST")
	}

	if config.database.Port = os.Getenv("DBPORT"); len(config.database.Host) == 0 {
		log.Fatal("No value in env PORT")
	}

	if config.database.Username = os.Getenv("DBUSER"); len(config.database.Host) == 0 {
		log.Fatal("No value in env DBUSER ")
	}

	if config.database.Password = os.Getenv("DBPASS"); len(config.database.Host) == 0 {
		log.Fatal("No value in env DBPASS ")
	}

	if config.database.Schema = os.Getenv("DATABASE"); len(config.database.Host) == 0 {
		log.Fatal("No value in env DATABASE ")
	}

	if config.service.Name = os.Getenv("APP_NAME"); len(config.database.Host) == 0 {
		log.Fatal("No value in env APP_NAME ")
	}

	if config.service.Port = os.Getenv("SERVICEPORT"); len(config.database.Host) == 0 {
		log.Fatal("No value in env APP_NAME ")
	}

	if config.service.Servicehost = os.Getenv("SERVICEHOST"); len(config.database.Host) == 0 {
		log.Fatal("No value in env APP_NAME ")
	}

	return &config
}

func (config *Configuration) GetService() *Service {
	return &config.service
}

func (config *Configuration) GetDatabase() *Database {
	return &config.database
}
