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
			Pattern:     "/users/register",
			HandlerFunc: r.c.Register,
		},
		{
			Method:      "POST",
			Pattern:     "/users/login",
			HandlerFunc: r.c.Login,
		},
		{
			Method:      "GET",
			Pattern:     "/users/login-check",
			HandlerFunc: r.c.LoginCheck,
		},
	}
}
