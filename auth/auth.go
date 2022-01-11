package auth

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/snowmerak/dakuaz"
	"github.com/snowmerak/rest-make/cookie"
	"github.com/snowmerak/rest-make/cookie/jwthmac"
	"github.com/valyala/fasthttp"
)

const token = "token"

var secret string

func init() {
	sec := os.Getenv("JWT_SECRET")
	if sec == "" {
		secret = "1q!2w@3e#4r$5t%6y^7u&8i*9o(0p)_+qazxswedc"
	}
	secret = sec
}

func JWT(ctx *fasthttp.RequestCtx, criteria func(jwt.MapClaims) bool) error {
	ck := ctx.Request.Header.Cookie(token)
	if ck == nil {
		ctx.Response.Header.SetStatusCode(fasthttp.StatusUnauthorized)
		return fmt.Errorf("no token")
	}
	claims, err := jwthmac.Deserialize(secret, string(ck))
	if err != nil {
		ctx.Response.Header.SetStatusCode(fasthttp.StatusUnauthorized)
		return fmt.Errorf("auth.JWT: %w", err)
	}
	if !criteria(claims) {
		ctx.Response.Header.SetStatusCode(fasthttp.StatusForbidden)
		return fmt.Errorf("token is disqualified")
	}
	ctx.Response.Header.SetStatusCode(fasthttp.StatusOK)
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
	if d.Verify(secret) != nil {
		ctx.Response.Header.SetStatusCode(fasthttp.StatusUnauthorized)
		return fmt.Errorf("token is disqualified")
	}
	if !criteria(d) {
		ctx.Response.Header.SetStatusCode(fasthttp.StatusForbidden)
		return fmt.Errorf("token is disqualified")
	}
	ctx.Response.Header.SetStatusCode(fasthttp.StatusOK)
	return nil
}

func MakeJWT(ctx *fasthttp.RequestCtx, claims jwt.StandardClaims, additional jwt.MapClaims) error {
	tk, err := jwthmac.Serialize(secret, claims, additional)
	if err != nil {
		return fmt.Errorf("auth.MakeJWT: %w", err)
	}
	ck := cookie.New(token, tk, time.Unix(claims.ExpiresAt, 0).UTC())
	ctx.Response.Header.SetCookie(ck)
	return nil
}

func MakeDakuaz(ctx *fasthttp.RequestCtx, id string, duration time.Duration) error {
	tk, err := dakuaz.New(secret, id, duration)
	if err != nil {
		return fmt.Errorf("auth.MakeDakuaz: %w", err)
	}
	ck := cookie.New(token, tk.Serialize(), time.Now().Add(duration))
	ctx.Response.Header.SetCookie(ck)
	return nil
}
