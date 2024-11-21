package utils

import (
	"context"

	"github.com/mcgtrt/go-puerto/types"
)

func GetLocale(ctx context.Context) (language string, currency string) {
	language, _ = ctx.Value(types.LanguageCtxKey{}).(string)
	currency, _ = ctx.Value(types.CurrencyCtxKey{}).(string)
	return
}

func Ptr[T any](v T) *T {
	return &v
}
