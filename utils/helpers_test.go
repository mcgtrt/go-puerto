package utils

import (
	"context"
	"testing"

	"github.com/mcgtrt/go-puerto/types"
	"github.com/stretchr/testify/assert"
)

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
