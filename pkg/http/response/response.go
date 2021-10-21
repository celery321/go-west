package response

import (
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/transport/http"
	stdHttp "net/http"
	"strings"
)

const (
	baseContentType = "application"
)

type Response struct {
	Code    int         `json:"code" form:"code"`
	Message string      `json:"message" form:"message"`
	Data    interface{} `json:"data" form:"data"`
}

func New()  *Response{
	return &Response{}
}

// ContentType returns the content-type with base prefix.
func ContentType(subtype string) string {
	return strings.Join([]string{baseContentType, subtype}, "/")
}

func ResponseEncoder(w stdHttp.ResponseWriter, r *stdHttp.Request, v interface{}) error {
	reply := New()
	reply.Code = 200
	reply.Data = v
	reply.Message = ""

	codec, _ := http.CodecForRequest(r, "Accept")
	data, err := codec.Marshal(reply)
	if err != nil {
		return err
	}
	w.Header().Set("Content-Type", ContentType(codec.Name()))
	w.WriteHeader(stdHttp.StatusOK)
	w.Write(data)
	return nil
}

func ErrorEncoder(w stdHttp.ResponseWriter, r *stdHttp.Request, err error) {
	se := errors.FromError(err)
	reply := New()
	reply.Code = int(se.Code)
	reply.Data = nil
	reply.Message = se.Message

	codec, _ := http.CodecForRequest(r, "Accept")
	body, err := codec.Marshal(reply)
	if err != nil {
		w.WriteHeader(stdHttp.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", ContentType(codec.Name()))
	w.WriteHeader(int(se.Code))
	w.Write(body)
}