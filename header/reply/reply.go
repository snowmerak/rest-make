package reply

import (
	"encoding/json"

	"github.com/snowmerak/rest-make/header/contenttype"
	"github.com/valyala/fasthttp"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

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
	ctx.SetContentType(string(contenttype.ApplicationProto))
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
