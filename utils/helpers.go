package utils

import (
	"context"
	"regexp"

	"github.com/mcgtrt/go-puerto/types"
)

// Check if given string is URL safe
func IsURLSafe(input string) bool {
	urlSafePattern := `^[a-zA-Z0-9._~:/?#\[\]@!$&'()*+,;=%-]*$`
	matched, err := regexp.MatchString(urlSafePattern, input)
	if err != nil {
		return false
	}
	return matched
}

// Helper function to get both language and currency from the context
func GetLocale(ctx context.Context) (language string, currency string) {
	language, _ = ctx.Value(types.LanguageCtxKey{}).(string)
	currency, _ = ctx.Value(types.CurrencyCtxKey{}).(string)
	return
}

// Return pointer of the value
func Ptr[T any](v T) *T {
	return &v
}
