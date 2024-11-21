package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEmails(t *testing.T) {
	tests := []struct {
		name    string
		email   string
		isValid bool
	}{
		// Valid Email Addresses
		{"Standard valid email", "email@test.com", true},
		{"Valid email with alias", "email+alias@test.com", true},
		{"Valid email with subdomain", "email.email@test.co.uk", true},
		{"Single character local part", "e@test.com", true},
		{"Numeric local part", "1234567890@test.com", true},
		{"Valid email with uncommon TLD", "email@test.travel", true},
		{"Local part with underscores", "_______@test.com", true},
		{"Valid email with uncommon TLD", "email@test.name", true},
		{"Local part with hyphen", "firstname-lastname@test.com", true},
		{"Email with subdomain", "email@sub.test.com", true},
		{"Email with nested subdomains", "e.mail@sub.test.travel", true},

		// Invalid Email Addresses
		{"Missing @", "plainaddress", false},
		{"Missing local part", "@test.com", false},
		{"Missing @ symbol", "email.test.com", false},
		{"Leading dot in domain", "email@.test.com", false},
		{"Missing top-level domain", "email@com", false},
		{"Double dots in domain", "email@test..com", false},
		{"Trailing dot in local part", "email.@test.com", false},
		{"Leading dot in local part", ".email@test.com", false},
		{"Leading hyphen in domain", "email@-test.com", false},
		{"Trailing hyphen in domain", "email@test.com-", false},
		{"Invalid IP address in domain", "email@111.222.333.44444", false},
		{"Invalid underscores in domain", "email@domain_with_underscores.com", false},
		{"Missing domain name", "email@.com", false},
		{"Invalid domain with numeric TLD", "email@.123.com", false},
		{"Double dots in domain", "email@domain..com", false},
		{"Multiple @ symbols", "email@domain@domain.com", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			valid := IsEmailCorrect(tt.email)
			assert.Equal(t, tt.isValid, valid, "valid check failed")
		})
	}
}

func TestIsNameCorrect(t *testing.T) {
	// Ensure MIN_NAME_LEN < MAX_NAME_LEN
	assert.Less(t, MIN_NAME_LEN, MAX_NAME_LEN, "MIN_NAME_LEN must be less than MAX_NAME_LEN")

	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{"Empty string", "", MIN_NAME_LEN <= 0},
		{"Below minimum length", "A", len("A") >= MIN_NAME_LEN},
		{"Exact minimum length", "AB", len("AB") >= MIN_NAME_LEN},
		{"Above minimum but below maximum", "ValidName", len("ValidName") >= MIN_NAME_LEN && len("ValidName") <= MAX_NAME_LEN},
		{"Exact maximum length", string(make([]byte, MAX_NAME_LEN)), len(string(make([]byte, MAX_NAME_LEN))) <= MAX_NAME_LEN},
		{"Above maximum length", string(make([]byte, MAX_NAME_LEN+1)), false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsNameCorrect(tt.input)
			assert.Equal(t, tt.expected, result, "For input: %s", tt.input)
		})
	}
}

func TestIsPasswordCorrect(t *testing.T) {
	// Ensure MIN_PASSWORD_LEN < MAX_PASSWORD_LEN
	assert.Less(t, MIN_PASSWORD_LEN, MAX_PASSWORD_LEN, "MIN_PASSWORD_LEN must be less than MAX_PASSWORD_LEN")

	tests := []struct {
		name     string
		password string
		expected bool
	}{
		{"Empty password", "", false},
		{"Below minimum length", "Short1!", false},
		{"Exact minimum length - invalid", "Short1", false},
		{"Exact minimum length - valid", "Valid11!", true},
		{"Above maximum length", string(make([]byte, MAX_PASSWORD_LEN+1)), false},
		{"Exact maximum length - valid", "ValidPassword123!@#", true},
		{"No uppercase letters", "password123!", false},
		{"No numbers", "Password!", false},
		{"No special characters", "Password123", false},
		{"Valid password", "Passw0rd!", true},
		{"Only special characters", "!@#$%^&*", false},
		{"Only numbers", "12345678", false},
		{"Only uppercase", "ABCDEFGH", false},
		{"Negative MIN_PASSWORD_LEN", "", MIN_PASSWORD_LEN <= 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsPasswordCorrect(tt.password)
			assert.Equal(t, tt.expected, result, "For password: %s", tt.password)
		})
	}
}
