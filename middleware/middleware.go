package middleware

import "github.com/valyala/fasthttp"

func New(funs ...func(ctx *fasthttp.RequestCtx) bool) func(func(ctx *fasthttp.RequestCtx)) func(ctx *fasthttp.RequestCtx) {
	return func(handler func(ctx *fasthttp.RequestCtx)) func(*fasthttp.RequestCtx) {
		return func(ctx *fasthttp.RequestCtx) {
			for _, fun := range funs {
				if !fun(ctx) {
					return
				}
			}
			handler(ctx)
		}
	}
}

var None = New()
