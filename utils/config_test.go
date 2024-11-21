package utils

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDefaultConfig(t *testing.T) {
	var (
		fileserverPath = "FILE_SERVER_PATH"
		httpPort       = "HTTP_PORT"
		importAlpine   = "IMPORT_ALPINE_JS"
		incMongo       = "INCLUDE_MONGO"
		incPostgres    = "INCLUDE_POSTGRES"
		incValkey      = "INCLUDE_VALKEY"
		mongodbName    = "MONGO_DB_NAME"
		mongoUsername  = "MONGO_USERNAME"
		mongoPassword  = "MONGO_PASSWORD"
		mongoHost      = "MONGO_HOST"
		mongoPort      = "MONGO_PORT"
	)

	defer os.Unsetenv(fileserverPath)
	defer os.Unsetenv(httpPort)
	defer os.Unsetenv(importAlpine)
	defer os.Unsetenv(incMongo)
	defer os.Unsetenv(incPostgres)
	defer os.Unsetenv(incValkey)
	defer os.Unsetenv(mongodbName)
	defer os.Unsetenv(mongoUsername)
	defer os.Unsetenv(mongoPassword)
	defer os.Unsetenv(mongoHost)
	defer os.Unsetenv(mongoPort)

	c, err := NewDefaultConfig()
	errInvalidPortString := "http port must be a valid port number"
	assertConfigNil(t, c, err, errInvalidPortString)

	os.Setenv(httpPort, "invalid")
	c, err = NewDefaultConfig()
	assertConfigNil(t, c, err, errInvalidPortString)

	errInvalidPortRange := "only registered and dynamic ports are allowed (1000 - 65535)"
	os.Setenv(httpPort, "15")
	c, err = NewDefaultConfig()
	assertConfigNil(t, c, err, errInvalidPortRange)

	os.Setenv(httpPort, "1512312")
	c, err = NewDefaultConfig()
	assertConfigNil(t, c, err, errInvalidPortRange)

	os.Setenv(httpPort, "3000")
	// os.Setenv(importAlpine)
	c, err = NewDefaultConfig()
	assert.Equal(t, 3000, c.HTTP.Port, "expected the same http port")
}

func assertConfigNil(t *testing.T, c *Config, err error, msg string) {
	assert.Nil(t, c, "expected nil config")
	assert.Error(t, err, "expected an error")
	assert.Equal(t, err.Error(), msg, "expected the same error message")
}
