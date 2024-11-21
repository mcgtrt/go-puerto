package utils

import (
	"regexp"
	"strings"
	"unicode"
)

const (
	MIN_NAME_LEN     = 2
	MIN_PASSWORD_LEN = 8

	MAX_NAME_LEN     = 64
	MAX_PASSWORD_LEN = 32
)

func IsEmailCorrect(email string) bool {
	regex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9-]+(\.[a-zA-Z0-9-]+)*\.[a-zA-Z]{2,10}$`
	matched, _ := regexp.MatchString(regex, email)
	if !matched {
		return false
	}

	// Additional programmatic checks
	parts := strings.Split(email, "@")
	local, domain := parts[0], parts[1]

	// Ensure no leading or trailing dots in the local part
	if strings.HasPrefix(local, ".") || strings.HasSuffix(local, ".") {
		return false
	}

	// Ensure no leading or trailing hyphens in the domain part
	if strings.HasPrefix(domain, "-") || strings.HasSuffix(domain, "-") {
		return false
	}

	// Ensure no consecutive dots in the domain part
	if strings.Contains(domain, "..") {
		return false
	}

	return true
}

func IsNameCorrect(name string) bool {
	if MIN_NAME_LEN > 0 {
		if len(name) < MIN_NAME_LEN {
			return false
		}
	}
	if len(name) > MAX_NAME_LEN {
		return false
	}
	return true
}

func IsPasswordCorrect(pwd string) bool {
	var number, upper, special = false, false, false

	if len(pwd) < MIN_PASSWORD_LEN || len(pwd) > MAX_PASSWORD_LEN {
		return false
	}

	for _, c := range pwd {
		switch {
		case unicode.IsNumber(c):
			number = true
		case unicode.IsUpper(c):
			upper = true
		case unicode.IsPunct(c) || unicode.IsSymbol(c):
			special = true
		}
	}

	return number && upper && special
}
