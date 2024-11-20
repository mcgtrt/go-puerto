package utils

import (
	"context"
	"errors"
	"os"
	"strconv"

	_ "github.com/joho/godotenv/autoload"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Holds global configuration sourced from local .env file
type Config struct {
	HTTP     *HTTPConfig
	Mongo    *MongoConfig
	Postgres *PostgresConfig
	Valkey   *ValkeyConfig
}

// Create new default config from the local .env file. If any part of the configuration
// will cause an error, will return an empty config (nil) and error message
func NewDefaultConfig() (*Config, error) {
	config := &Config{}

	http, err := newDefaultHTTPConfig()
	if err != nil {
		return nil, err
	}
	config.HTTP = http

	if os.Getenv("INCLUDE_MONGO") == "true" {
		mongo, err := newDefaultMongoConfig()
		if err != nil {
			return nil, err
		}
		config.Mongo = mongo
	}
	if os.Getenv("INCLUDE_POSTGRES") == "true" {
		postgres, err := newDefaultPostgresConfig()
		if err != nil {
			return nil, err
		}
		config.Postgres = postgres
	}
	if os.Getenv("INCLUDE_VALKEY") == "true" {
		valkey, err := newDefaultValkeyConfig()
		if err != nil {
			return nil, err
		}
		config.Valkey = valkey
	}

	return config, nil
}

// Configuration required for HTTP server
type HTTPConfig struct {
	Port int

	WithHTMX     bool
	WithAlpineJS bool
}

func newDefaultHTTPConfig() (*HTTPConfig, error) {
	config := &HTTPConfig{}
	port, err := strconv.Atoi(os.Getenv("HTTP_PORT"))
	if err != nil {
		return nil, errors.New("http port must be a valid port number")
	}
	config.Port = port
	if os.Getenv("INCLUDE_HTMX") == "true" {
		config.WithHTMX = true
	}
	if os.Getenv("INCLUDE_ALPINE_JS") == "true" {
		config.WithAlpineJS = true
	}
	return config, nil
}

// Required configuration for creating mongodb connection
type MongoConfig struct {
	Client *mongo.Client
	DBName string
}

func newDefaultMongoConfig() (*MongoConfig, error) {
	dbname := os.Getenv("MONGO_DB_NAME")
	if dbname == "" {
		return nil, errors.New("mongo database name cannot be empty")
	}

	var (
		username    = os.Getenv("MONGO_USERNAME")
		password    = os.Getenv("MONGO_PASSWORD")
		host        = os.Getenv("MONGO_HOST")
		port        = os.Getenv("MONGO_PORT")
		dburl       = "mongodb://" + username + ":" + password + "@" + host + ":" + port
		client, err = mongo.Connect(context.Background(), options.Client().ApplyURI(dburl))
	)
	if err != nil {
		return nil, err
	}

	return &MongoConfig{
		Client: client,
		DBName: dbname,
	}, nil
}

// Required configuration for creating postgres connection
type PostgresConfig struct{}

func newDefaultPostgresConfig() (*PostgresConfig, error) {
	return &PostgresConfig{}, nil
}

// Required configuration for creating valkey connection
type ValkeyConfig struct{}

func newDefaultValkeyConfig() (*ValkeyConfig, error) {
	return &ValkeyConfig{}, nil
}
