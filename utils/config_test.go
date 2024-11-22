package utils

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDefaultConfig(t *testing.T) {
	for _, key := range AllConfigKeys() {
		defer os.Unsetenv(key)
	}
	// Test HTTP Config
	ts := "true"
	errInvalidPortString := "http port must be a valid port number"
	errInvalidPortRange := "only registered and dynamic ports are allowed (1000 - 65535)"
	httptests := []struct {
		name string
		key  string
		val  string
		err  string
	}{
		{
			"start",
			"",
			"",
			errInvalidPortString,
		},
		{
			"invalid http port",
			HTTP_PORT,
			"invalid",
			errInvalidPortString,
		},
		{
			"invalid port range - too little",
			HTTP_PORT,
			"15",
			errInvalidPortRange,
		},
		{
			"invalid port range - too big",
			HTTP_PORT,
			"1512312",
			errInvalidPortRange,
		},
	}
	for _, tt := range httptests {
		if tt.key != "" {
			os.Setenv(tt.key, tt.val)
			defer os.Unsetenv(tt.key)
		}
		c, err := NewDefaultConfig()
		assertConfigNil(t, c, err, tt.err)
	}

	os.Setenv(HTTP_PORT, "3000")
	c, err := NewDefaultConfig()
	assert.Equal(t, 3000, c.HTTP.Port, "expected the same http port")
	assert.Nil(t, err, "expected no errors")

	os.Setenv(FILE_SERVER_PATH, "non url safe path")
	c, err = NewDefaultConfig()
	assertConfigNil(t, c, err, "file server path is not URL safe")

	os.Setenv(FILE_SERVER_PATH, "static")
	os.Setenv(USE_JS_ALPINE, ts)
	c, err = NewDefaultConfig()
	assert.True(t, c.HTTP.ImportAlpineJS, "should be true")
	assert.Equal(t, "static", c.HTTP.FileServerPath)
	assert.Nil(t, err, "expected no errors")
	assert.Nil(t, c.Mongo, "expected nil mongo config")
	assert.Nil(t, c.Postgres, "expected nil postgres config")
	assert.Nil(t, c.Valkey, "expected nil valkey config")

	// Test Middleware Config
	mws := []string{
		USE_MW_LOCALISATION,
		USE_MW_SECURE_HEADERS,
		USE_MW_RATE_LIMIT,
		USE_MW_LOG_AND_MONITOR_HEADERS,
		USE_MW_CORS,
		USE_MW_ETAG,
		USE_MW_VALIDATE_SANITISE_HEADERS,
		USE_MW_METHOD_OVERRIDE,
	}
	for _, mw := range mws {
		os.Setenv(mw, ts)
	}
	c, err = NewDefaultConfig()
	assert.True(t, c.Middleware.Localisation, "expected localisation mw true")
	assert.True(t, c.Middleware.SecureHeaders, "expected secure headers mw true")
	assert.True(t, c.Middleware.RateLimit, "expected rate limit mw true")
	assert.True(t, c.Middleware.LogAndMonitorHeaders, "expected log and monitor mw true")
	assert.True(t, c.Middleware.CORS, "expected cors mw true")
	assert.True(t, c.Middleware.ETAG, "expected etag mw true")
	assert.True(t, c.Middleware.ValidateSanitiseHeaders, "expected validate and sanitise headers mw true")
	assert.True(t, c.Middleware.MethodOverride, "expected method override mw true")
	assert.Nil(t, err, "expected no errors")

	os.Setenv(MW_RATE_LIMITER_LIMIT, "invalid")
	os.Setenv(MW_RATE_LIMITER_BURST, "invalidtoo")
	c, err = NewDefaultConfig()
	assertConfigNil(t, c, err, "rate limiter limit must be a number")

	os.Setenv(MW_RATE_LIMITER_LIMIT, "50")
	c, err = NewDefaultConfig()
	assertConfigNil(t, c, err, "rate limiter burst must be a number")

	os.Setenv(MW_RATE_LIMITER_BURST, "10")
	c, err = NewDefaultConfig()
	assertConfigNil(t, c, err, "rate limiter limit cannot be bigger than limiter burst")

	os.Setenv(MW_RATE_LIMITER_BURST, "100")
	c, err = NewDefaultConfig()
	assert.Equal(t, 50, *c.Middleware.RateLimiterLimit, "expected the same rate limiter limit")
	assert.Equal(t, 100, *c.Middleware.RateLimiterBurst, "expected the same rate limiter burst")
	assert.Nil(t, err, "expected no errors")

	// Test Mongo Config
	os.Setenv(USE_DB_MONGO, "invalid")
	c, err = NewDefaultConfig()
	assert.Nil(t, c.Mongo, "expected nil mongo config")
	assert.Nil(t, err, "expected no errors")

	os.Setenv(USE_DB_MONGO, ts)
	c, err = NewDefaultConfig()
	assertConfigNil(t, c, err, "mongo database name cannot be empty")

	os.Setenv(MONGO_DB_NAME, "testname")
	c, err = NewDefaultConfig()
	assertConfigNil(t, c, err, "mongo username cannot be empty")

	os.Setenv(MONGO_USERNAME, "testusername")
	c, err = NewDefaultConfig()
	assertConfigNil(t, c, err, "mongo password cannot be empty")

	os.Setenv(MONGO_PASSWORD, "password")
	c, err = NewDefaultConfig()
	assert.Equal(t, "testname", c.Mongo.DBName, "expected the same mongo database name")
	assert.Equal(t, "testusername", c.Mongo.Username, "expected the same usernames")
	assert.Equal(t, "password", c.Mongo.Password, "expected the same passwords")
	assert.Equal(t, "27017", c.Mongo.Port, "expected the same mongo ports")
	assert.Equal(t, "localhost", c.Mongo.Host, "expected the same hosts")
	assert.Nil(t, err, "expected no errors")

	os.Setenv(MONGO_PORT, "invalid port")
	c, err = NewDefaultConfig()
	assertConfigNil(t, c, err, "invalid port for mongo connection")

	os.Setenv(MONGO_PORT, "123456")
	os.Setenv(MONGO_HOST, "someotherhost")
	c, err = NewDefaultConfig()
	assert.Equal(t, "123456", c.Mongo.Port, "expected the same ports")
	assert.Equal(t, "someotherhost", c.Mongo.Host, "expected the same hosts")
	assert.Nil(t, err, "expected no errors")

	// Test Postgres Config
	os.Setenv(USE_DB_POSTGRES, "invalid")
	c, err = NewDefaultConfig()
	assert.Nil(t, c.Postgres, "expected nil postgres config")
	assert.Nil(t, err, "expected no errors")

	os.Setenv(USE_DB_POSTGRES, ts)
	c, err = NewDefaultConfig()
	assert.NotNil(t, c.Postgres, "expected postgres config")
	assert.Nil(t, err, "expected no errors")

	// Test Valkey Config
	os.Setenv(USE_DB_VALKEY, "invalid")
	c, err = NewDefaultConfig()
	assert.Nil(t, c.Valkey, "expected nil valkey config")
	assert.Nil(t, err, "expected no errors")

	os.Setenv(USE_DB_VALKEY, ts)
	c, err = NewDefaultConfig()
	assert.NotNil(t, c.Valkey, "expected valkey cnofiguration")
	assert.Nil(t, err, "expected no errors")
}

func assertConfigNil(t *testing.T, c *Config, err error, msg string) {
	assert.Nil(t, c, "expected nil config")
	assert.Error(t, err, "expected an error")
	assert.Equal(t, err.Error(), msg, "expected the same error message")
}
