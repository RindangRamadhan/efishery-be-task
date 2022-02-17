package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/rindangramadhan/efishery-be-task/internal/helper"
	"github.com/rindangramadhan/efishery-be-task/internal/modules/health"
)

type ApiRoute interface {
	Routes() []helper.Route
}

func RouteApply(e *echo.Echo) {
	var routers []helper.Route

	apiRoute := []ApiRoute{
		health.NewRoute(),
	}

	for _, ar := range apiRoute {
		routers = append(routers, ar.Routes()...)
	}

	api := e.Group("api")
	for _, router := range routers {

		switch router.Method {
		case echo.GET:
			{
				api.GET(router.Pattern, router.HandlerFunc, router.Middleware...)
				break
			}
		case echo.POST:
			{
				api.POST(router.Pattern, router.HandlerFunc, router.Middleware...)
				break
			}
		}

	}
}
