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
// @Tags         Health
// @Description  To check health service
// @Accept       json
// @Produce      json
// @Success      200  {object}  response.BodyTpl{status=string,message=string,values=response.Object}
// @Failure      500  {object}  response.BodyTpl{status=string,message=string,errors=response.Object}
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
