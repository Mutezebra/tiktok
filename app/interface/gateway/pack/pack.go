package pack

import (
	"errors"
	"net/http"

	"github.com/Mutezebra/tiktok/pkg/log"

	"github.com/Mutezebra/tiktok/app/domain/model/errno"

	"github.com/cloudwego/hertz/pkg/app"
)

type Response struct {
	Base Base `json:"base"`
	Data any  `json:"data,omitempty"`
}

type Base struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func SendResponse(c *app.RequestContext, data any) {
	resp := &Response{
		Base: Base{
			200,
			"operate success",
		},
		Data: data,
	}

	c.JSON(http.StatusOK, resp)
}

func SendFailedResponse(c *app.RequestContext, error error) {
	log.LogrusObj.Error(error)

	var e errno.Errno
	ok := errors.As(error, &e)
	if !ok {
		e = errno.Convert(error)
	}

	resp := &Response{
		Base: Base{
			e.Code(),
			e.Error(),
		},
	}
	c.JSON(http.StatusOK, resp)
}
