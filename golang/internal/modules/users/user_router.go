package users

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
			Method:      "POST",
			Pattern:     "/auth/register",
			HandlerFunc: r.c.Register,
		},
		{
			Method:      "POST",
			Pattern:     "/auth/login",
			HandlerFunc: r.c.Login,
		},
		{
			Method:      "GET",
			Pattern:     "/auth/login-check",
			HandlerFunc: r.c.LoginCheck,
		},
	}
}
