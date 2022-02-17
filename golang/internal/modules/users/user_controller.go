package users

import (
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/rindangramadhan/efishery-be-task/pkg/response"

	upkg "github.com/rindangramadhan/efishery-be-task/pkg/request/users"
)

type Controller interface {
	Register(ctx echo.Context) (err error)
}

type controller struct {
	m Model
}

func NewController() Controller {
	return &controller{
		m: NewModel(),
	}
}

func (c *controller) Register(ctx echo.Context) (err error) {
	// Bind & Validation Request
	req := new(upkg.RegisterRequest)
	if err := ctx.Bind(req); err != nil {
		return response.WriteError(ctx, http.StatusBadRequest, http.StatusText(http.StatusBadRequest), err)
	}

	if err := ctx.Validate(req); err != nil {
		return response.WriteError(ctx, http.StatusBadRequest, http.StatusText(http.StatusBadRequest), err)
	}

	resp, err := c.m.Register(ctx, req)

	if err != nil {
		if strings.Contains(err.Error(), "UNIQUE constraint failed: users.name") {
			stackEntries := map[string]interface{}{
				"message": []interface{}{
					map[string]interface{}{
						"field":   "name",
						"message": "name already exist",
					},
				},
			}

			return response.WriteError(ctx, http.StatusConflict, http.StatusText(http.StatusConflict), stackEntries)
		} else {
			return response.WriteError(ctx, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), err)
		}
	}

	return response.WriteSuccess(ctx, "Successfully Register", resp)
}
