package utils

import (
	"context"
	"testing"

	"github.com/mcgtrt/go-puerto/types"
	"github.com/stretchr/testify/assert"
)

func TestIsURLSafe(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		// Valid cases
		{"Simple alphanumeric", "hello-world123", true},
		{"Unreserved characters", "._~", true},
		{"Reserved characters", ":/?#[]@!$&'()*+,;=", true},
		{"Full URL", "https://example.com/path?query=value#fragment", true},
		{"Email URL", "mailto:example@test.com", true},
		{"IPv6 Address", "[IPv6:2001:db8::1]", true},
		{"Query string with special characters", "?key=value&foo=bar", true},
		{"Fragment identifier", "#fragment", true},
		{"Mixed valid characters", "https://user:pass@example.com:8080/path?query#frag", true},
		{"Percent encoding", "%20%3A%2F%3F", true},
		{"Empty string", "", true},

		// Invalid cases
		{"Contains pipe character", "example|test", false},
		{"Contains angle brackets", "example<test>", false},
		{"Contains curly braces", "example{test}", false},
		{"Contains double quotes", `example"test"`, false},
		{"Contains backslash", `example\test`, false},
		{"Contains caret", "example^test", false},
		{"Contains backtick", "example`test", false},
		{"Contains space", "example test", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsURLSafe(tt.input)
			assert.Equal(t, tt.expected, result, "Input: %s", tt.input)
		})
	}
}

func TestGetLocale(t *testing.T) {
	ctx := context.Background()
	lang, curr := GetLocale(ctx)
	assert.Empty(t, lang, "expected language to be empty")
	assert.Empty(t, curr, "expected currency to be empty")

	l := "en"
	c := "GBP"
	ctx = context.WithValue(ctx, types.LanguageCtxKey{}, l)
	ctx = context.WithValue(ctx, types.CurrencyCtxKey{}, c)
	lang, curr = GetLocale(ctx)
	assert.Equal(t, l, lang, "expected the same languages")
	assert.Equal(t, c, curr, "expected the same currencies")
}

func TestPtr(t *testing.T) {
	s := "test"
	i := 0
	sptr := Ptr(s)
	iptr := Ptr(i)
	assert.Equal(t, s, *sptr, "expected the same string values")
	assert.Equal(t, i, *iptr, "expected the same integer values")
}
