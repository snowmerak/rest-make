package header

import "github.com/valyala/fasthttp"

func SetContentType(ctx *fasthttp.RequestCtx, contentType string) {
	ctx.Response.Header.SetContentType(contentType)
}

func SetCookie(ctx *fasthttp.RequestCtx, cookie *fasthttp.Cookie) {
	ctx.Response.Header.SetCookie(cookie)
}

func GetCookie(ctx *fasthttp.RequestCtx, key string) []byte {
	return ctx.Request.Header.Cookie(key)
}
