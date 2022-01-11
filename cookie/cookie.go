package cookie

import (
	"time"

	"github.com/valyala/fasthttp"
)

func New(key string, value string, expireTime time.Time) *fasthttp.Cookie {
	ck := fasthttp.AcquireCookie()
	ck.SetHTTPOnly(true)
	ck.SetSecure(true)
	ck.SetExpire(expireTime)
	ck.SetKey(key)
	ck.SetValue(value)
	return ck
}
