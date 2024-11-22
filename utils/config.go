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

const (
	PROJECT_NAME                     = "PROJECT_NAME"
	FILE_SERVER_PATH                 = "FILE_SERVER_PATH"
	AES_SECRET                       = "AES_SECRET"
	USE_DB_MONGO                     = "USE_DB_MONGO"
	USE_DB_POSTGRES                  = "USE_DB_POSTGRES"
	USE_DB_VALKEY                    = "USE_DB_VALKEY"
	USE_JS_ALPINE                    = "USE_JS_ALPINE"
	USE_MW_LOCALISATION              = "USE_MW_LOCALISATION"
	USE_MW_SECURE_HEADERS            = "USE_MW_SECURE_HEADERS"
	USE_MW_RATE_LIMIT                = "USE_MW_RATE_LIMIT"
	MW_RATE_LIMITER_LIMIT            = "MW_RATE_LIMITER_LIMIT"
	MW_RATE_LIMITER_BURST            = "MW_RATE_LIMITER_BURST"
	USE_MW_LOG_AND_MONITOR_HEADERS   = "USE_MW_LOG_AND_MONITOR_HEADERS"
	USE_MW_CORS                      = "USE_MW_CORS"
	USE_MW_ETAG                      = "USE_MW_ETAG"
	USE_MW_VALIDATE_SANITISE_HEADERS = "USE_MW_VALIDATE_SANITISE_HEADERS"
	USE_MW_METHOD_OVERRIDE           = "USE_MW_METHOD_OVERRIDE"
	HTTP_PORT                        = "HTTP_PORT"
	MONGO_DB_NAME                    = "MONGO_DB_NAME"
	MONGO_USERNAME                   = "MONGO_USERNAME"
	MONGO_PASSWORD                   = "MONGO_PASSWORD"
	MONGO_HOST                       = "MONGO_HOST"
	MONGO_PORT                       = "MONGO_PORT"
)

func AllConfigKeys() []string {
	return []string{
		PROJECT_NAME,
		FILE_SERVER_PATH,
		AES_SECRET,
		USE_DB_MONGO,
		USE_DB_POSTGRES,
		USE_DB_VALKEY,
		USE_JS_ALPINE,
		USE_MW_LOCALISATION,
		USE_MW_SECURE_HEADERS,
		USE_MW_RATE_LIMIT,
		MW_RATE_LIMITER_LIMIT,
		MW_RATE_LIMITER_BURST,
		USE_MW_LOG_AND_MONITOR_HEADERS,
		USE_MW_CORS,
		USE_MW_ETAG,
		USE_MW_VALIDATE_SANITISE_HEADERS,
		USE_MW_METHOD_OVERRIDE,
		HTTP_PORT,
		MONGO_DB_NAME,
		MONGO_USERNAME,
		MONGO_PASSWORD,
		MONGO_HOST,
		MONGO_PORT,
	}
}

// Holds global configuration sourced from local .env file
type Config struct {
	HTTP       *HTTPConfig
	Middleware *MiddlewareConfig
	Mongo      *MongoConfig
	Postgres   *PostgresConfig
	Valkey     *ValkeyConfig
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
	mw, err := newDefaultMiddlewareConfig()
	if err != nil {
		return nil, err
	}
	config.Middleware = mw

	if os.Getenv(USE_DB_MONGO) == "true" {
		mongo, err := newDefaultMongoConfig()
		if err != nil {
			return nil, err
		}
		config.Mongo = mongo
	}
	if os.Getenv(USE_DB_POSTGRES) == "true" {
		postgres, err := newDefaultPostgresConfig()
		if err != nil {
			return nil, err
		}
		config.Postgres = postgres
	}
	if os.Getenv(USE_DB_VALKEY) == "true" {
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
	FileServerPath string
	ImportAlpineJS bool
	Port           int
}

func newDefaultHTTPConfig() (*HTTPConfig, error) {
	config := &HTTPConfig{}
	port, err := strconv.Atoi(os.Getenv(HTTP_PORT))
	if err != nil {
		return nil, errors.New("http port must be a valid port number")
	}
	if port < 1000 || port > 65535 {
		return nil, errors.New("only registered and dynamic ports are allowed (1000 - 65535)")
	}
	config.Port = port
	if os.Getenv(USE_JS_ALPINE) == "true" {
		config.ImportAlpineJS = true
	}
	path := os.Getenv(FILE_SERVER_PATH)
	if !IsURLSafe(path) {
		return nil, errors.New("file server path is not URL safe")
	}
	config.FileServerPath = path
	return config, nil
}

type MiddlewareConfig struct {
	Localisation            bool
	SecureHeaders           bool
	RateLimit               bool
	RateLimiterLimit        *int
	RateLimiterBurst        *int
	LogAndMonitorHeaders    bool
	CORS                    bool
	ETAG                    bool
	ValidateSanitiseHeaders bool
	MethodOverride          bool
}

func newDefaultMiddlewareConfig() (*MiddlewareConfig, error) {
	cfg := &MiddlewareConfig{}
	if loc := os.Getenv(USE_MW_LOCALISATION); loc == "true" {
		cfg.Localisation = true
	}
	if sec := os.Getenv(USE_MW_SECURE_HEADERS); sec == "true" {
		cfg.SecureHeaders = true
	}
	if rate := os.Getenv(USE_MW_RATE_LIMIT); rate == "true" {
		cfg.RateLimit = true
	}
	if limit := os.Getenv(MW_RATE_LIMITER_LIMIT); limit != "" {
		l, err := strconv.Atoi(limit)
		if err != nil {
			return nil, errors.New("rate limiter limit must be a number")
		}
		if l > 0 {
			cfg.RateLimiterLimit = &l
		}
	}
	if burst := os.Getenv(MW_RATE_LIMITER_BURST); burst != "" {
		b, err := strconv.Atoi(burst)
		if err != nil {
			return nil, errors.New("rate limiter burst must be a number")
		}
		if b > 0 {
			cfg.RateLimiterBurst = &b
		}
	}
	if cfg.RateLimiterLimit != nil && cfg.RateLimiterBurst != nil {
		if *cfg.RateLimiterLimit > *cfg.RateLimiterBurst {
			return nil, errors.New("rate limiter limit cannot be bigger than limiter burst")
		}
	}
	if log := os.Getenv(USE_MW_LOG_AND_MONITOR_HEADERS); log == "true" {
		cfg.LogAndMonitorHeaders = true
	}
	if cors := os.Getenv(USE_MW_CORS); cors == "true" {
		cfg.CORS = true
	}
	if etag := os.Getenv(USE_MW_ETAG); etag == "true" {
		cfg.ETAG = true
	}
	if val := os.Getenv(USE_MW_VALIDATE_SANITISE_HEADERS); val == "true" {
		cfg.ValidateSanitiseHeaders = true
	}
	if overr := os.Getenv(USE_MW_METHOD_OVERRIDE); overr == "true" {
		cfg.MethodOverride = true
	}
	return cfg, nil
}

// Required configuration for creating mongodb connection
type MongoConfig struct {
	Username string
	Password string
	Host     string
	Port     string
	DBName   string
}

func (c *MongoConfig) ConnectionString() string {
	if c.Username == "" || c.Password == "" || c.Host == "" || c.Port == "" {
		return ""
	}
	return "mongodb://" + c.Username + ":" + c.Password + "@" + c.Host + ":" + c.Port
}

func (c *MongoConfig) Client() (*mongo.Client, error) {
	ctx := context.Background()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(c.ConnectionString()))
	if err != nil {
		return nil, errors.New("error creating mongo client: " + err.Error())
	}
	if err := client.Ping(ctx, nil); err != nil {
		return nil, errors.New("error connecting to mongodb - could not ping the database: " + err.Error())
	}
	return client, nil
}

func newDefaultMongoConfig() (*MongoConfig, error) {
	dbname := os.Getenv(MONGO_DB_NAME)
	if dbname == "" {
		return nil, errors.New("mongo database name cannot be empty")
	}
	username := os.Getenv(MONGO_USERNAME)
	if username == "" {
		return nil, errors.New("mongo username cannot be empty")
	}
	password := os.Getenv(MONGO_PASSWORD)
	if password == "" {
		return nil, errors.New("mongo password cannot be empty")
	}
	host := os.Getenv(MONGO_HOST)
	if host == "" {
		host = "localhost"
	}
	port := os.Getenv(MONGO_PORT)
	if port == "" {
		port = "27017"
	}
	if _, err := strconv.Atoi(port); err != nil {
		return nil, errors.New("invalid port for mongo connection")
	}
	return &MongoConfig{
		DBName:   dbname,
		Username: username,
		Password: password,
		Host:     host,
		Port:     port,
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
