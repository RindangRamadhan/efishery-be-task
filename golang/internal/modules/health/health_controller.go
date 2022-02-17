package health

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Controller interface {
	Status(ctx echo.Context) error
}

type controller struct {
}

func NewController() Controller {
	return &controller{}
}

// Status godoc
// @Description  To check health service
// @Tags         Health
// @Accept       json
// @Produce      json
// @Success      200  {object}  response.BodyTpl{status=string,message=string}
// @Router       /health/status [get]
func (*controller) Status(ctx echo.Context) error {
	type Response struct {
		Status  string `json:"status"`
		Message string `json:"message"`
	}

	var response = Response{
		Status:  "success",
		Message: "Health check ok!",
	}

	return ctx.JSON(http.StatusOK, response)
}
