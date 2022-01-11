package auth

import (
	"fmt"

	"github.com/golang-jwt/jwt"
	"github.com/snowmerak/dakuaz"
	"github.com/snowmerak/rest-make/cookie/jwthmac"
	"github.com/valyala/fasthttp"
)

const token = "token"

func JWT(ctx *fasthttp.RequestCtx, criteria func(jwt.MapClaims) bool) error {
	ck := ctx.Request.Header.Cookie(token)
	if ck == nil {
		ctx.Response.Header.SetStatusCode(fasthttp.StatusUnauthorized)
		return fmt.Errorf("no token")
	}
	claims, err := jwthmac.Deserialize(string(ck))
	if err != nil {
		ctx.Response.Header.SetStatusCode(fasthttp.StatusUnauthorized)
		return fmt.Errorf("auth.JWT: %w", err)
	}
	if !criteria(claims) {
		ctx.Response.Header.SetStatusCode(fasthttp.StatusForbidden)
		return fmt.Errorf("token is disqualified")
	}
	return nil
}

func Dakuaz(ctx *fasthttp.RequestCtx, criteria func(*dakuaz.Dakuaz) bool) error {
	ck := ctx.Request.Header.Cookie(token)
	if ck == nil {
		ctx.Response.Header.SetStatusCode(fasthttp.StatusUnauthorized)
		return fmt.Errorf("no token")
	}
	d, err := dakuaz.Desirialize(string(ck))
	if err != nil {
		ctx.Response.Header.SetStatusCode(fasthttp.StatusUnauthorized)
		return fmt.Errorf("auth.Dakuaz: %w", err)
	}
	if !criteria(d) {
		ctx.Response.Header.SetStatusCode(fasthttp.StatusForbidden)
		return fmt.Errorf("token is disqualified")
	}
	return nil
}
