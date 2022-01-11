package reply

import (
	"encoding/json"

	"github.com/snowmerak/rest-make/header/contenttype"
	"github.com/valyala/fasthttp"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func Data(ctx *fasthttp.RequestCtx, contentType contenttype.ContentType, data []byte) error {
	ctx.SetContentType(string(contentType))
	if _, err := ctx.Write(data); err != nil {
		ctx.SetStatusCode(fasthttp.StatusServiceUnavailable)
		return err
	}
	ctx.SetStatusCode(fasthttp.StatusOK)
	return nil
}

func JSON(ctx *fasthttp.RequestCtx, data interface{}) error {
	ctx.SetContentType("application/json")
	encoder := json.NewEncoder(ctx)
	if err := encoder.Encode(data); err != nil {
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		return err
	}
	ctx.SetStatusCode(fasthttp.StatusOK)
	return nil
}

func Protobuf(ctx *fasthttp.RequestCtx, data protoreflect.ProtoMessage) error {
	ctx.SetContentType("application/protobuf")
	buf, err := proto.Marshal(data)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		return err
	}
	if _, err := ctx.Write(buf); err != nil {
		ctx.SetStatusCode(fasthttp.StatusServiceUnavailable)
		return err
	}
	ctx.SetStatusCode(fasthttp.StatusOK)
	return nil
}

func Plain(ctx *fasthttp.RequestCtx, data string) error {
	ctx.SetContentType("text/plain")
	if _, err := ctx.WriteString(data); err != nil {
		ctx.SetStatusCode(fasthttp.StatusServiceUnavailable)
		return err
	}
	ctx.SetStatusCode(fasthttp.StatusOK)
	return nil
}

func HTML(ctx *fasthttp.RequestCtx, data []byte) error {
	ctx.SetContentType("text/html")
	if _, err := ctx.Write(data); err != nil {
		ctx.SetStatusCode(fasthttp.StatusServiceUnavailable)
		return err
	}
	ctx.SetStatusCode(fasthttp.StatusOK)
	return nil
}

func Binary(ctx *fasthttp.RequestCtx, data []byte) error {
	ctx.SetContentType("application/octet-stream")
	if _, err := ctx.Write(data); err != nil {
		ctx.SetStatusCode(fasthttp.StatusServiceUnavailable)
		return err
	}
	ctx.SetStatusCode(fasthttp.StatusOK)
	return nil
}
