package utils

import (
	"context"

	"github.com/mcgtrt/go-puerto/api/middleware"
)

func GetLocale(ctx context.Context) (string, string) {
	lang, _ := ctx.Value(middleware.LocaleCtx{}).(string)
	currency, _ := ctx.Value(middleware.CurrencyCtx{}).(string)
	return lang, currency
}

func Ptr[T any](v T) *T {
	return &v
}
