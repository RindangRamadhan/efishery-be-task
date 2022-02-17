package health

import (
	"github.com/rindangramadhan/efishery-be-task/internal/helper"
)

type Route interface {
	Routes() []helper.Route
}

type route struct {
	c Controller
}

func NewRoute() Route {
	return &route{
		c: NewController(),
	}
}

func (r *route) Routes() []helper.Route {
	return []helper.Route{
		{
			Method:      "GET",
			Pattern:     "/health/status",
			HandlerFunc: r.c.Status,
		},
	}
}
