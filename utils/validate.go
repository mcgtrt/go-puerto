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
	var (
		emailRegex = regexp.MustCompile(`^[a-z0-9._%\$+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
		lower      = strings.ToLower(email)
	)
	return emailRegex.MatchString(lower)
}

func IsNameCorrect(name string) bool {
	if len(name) < MIN_NAME_LEN || len(name) > MAX_NAME_LEN {
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

	if !number || !upper || !special {
		return false
	}

	return true
}
