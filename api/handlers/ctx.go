package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/a-h/templ"
)

type Ctx struct {
	Context  context.Context
	Response http.ResponseWriter
	Request  *http.Request
}

func NewCtx(w http.ResponseWriter, r *http.Request) *Ctx {
	return &Ctx{
		Context:  r.Context(),
		Response: w,
		Request:  r,
	}
}

func (c *Ctx) Render(component templ.Component) error {
	return component.Render(c.Context, c.Response)
}

func (c *Ctx) JSON(code int, v any) error {
	c.Response.Header().Set("Content-Type", "application/json")
	c.Response.WriteHeader(code)
	return json.NewEncoder(c.Response).Encode(v)
}

func (c *Ctx) Text(code int, v string) error {
	c.Response.Header().Set("Content-Type", "text/plain; charset=utf-8")
	c.Response.WriteHeader(code)
	_, err := c.Response.Write([]byte(v))
	return err
}

func (c *Ctx) Error(code int) {
	http.Error(c.Response, http.StatusText(code), code)
}

func (c *Ctx) CloseBody() {
	c.Request.Body.Close()
}
