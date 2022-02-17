package routes

import (
	"os"

	"github.com/labstack/echo/v4"
	"github.com/rindangramadhan/efishery-be-task/docs"

	es "github.com/swaggo/echo-swagger"
)

type SwaggerInterface interface {
	Init(e *echo.Echo)
}

type swagger struct {
}

func (*swagger) Init(e *echo.Echo) {
	docs.SwaggerInfo_swagger.Title = os.Getenv("SWAGGER_TITLE")
	docs.SwaggerInfo_swagger.Description = os.Getenv("SWAGGER_DESCRIPTION")
	docs.SwaggerInfo_swagger.Version = os.Getenv("SWAGGER_VERSION")
	docs.SwaggerInfo_swagger.Host = os.Getenv("SWAGGER_HOST")
	docs.SwaggerInfo_swagger.BasePath = os.Getenv("SWAGGER_BASE_PATH")
	docs.SwaggerInfo_swagger.Schemes = []string{os.Getenv("SWAGGER_SCHEME")}

	api := e.Group("swagger")
	api.GET("/*", es.WrapHandler)
}

var Swagger = &swagger{}
