package response

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type (
	BodyTpl struct {
		Status  string      `json:"status"`
		Message string      `json:"message"`
		Values  interface{} `json:"values,omitempty"`
		Errors  interface{} `json:"errors,omitempty"`
	}

	Response struct {
		Code    string      `json:"code"`
		Status  string      `json:"status"`
		Message interface{} `json:"message"`
		Value   interface{} `json:"value"`
	}

	Object struct{}
)

func WriteSuccess(ctx echo.Context, message string, data interface{}) error {
	body := BodyTpl{
		Status:  "success",
		Message: message,
		Values:  data,
	}

	return ctx.JSON(http.StatusOK, body)
}

func WriteError(ctx echo.Context, httpStatusCode int, message string, stackEntries interface{}) error {
	body := BodyTpl{
		Status:  "error",
		Message: message,
		Errors:  stackEntries,
	}

	if stackEntries != nil {

		switch v := stackEntries.(type) {
		case *echo.HTTPError:
			body.Errors = v
		case error:
			body.Errors = v
		}

	}

	return ctx.JSON(httpStatusCode, body)
}
