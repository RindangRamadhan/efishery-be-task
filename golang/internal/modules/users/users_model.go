package users

import (
	"github.com/labstack/echo/v4"

	db "github.com/rindangramadhan/efishery-be-task/internal/db/sqlite"
	"github.com/rindangramadhan/efishery-be-task/internal/helper"
	upkg "github.com/rindangramadhan/efishery-be-task/pkg/request/users"
)

type Model interface {
	Register(c echo.Context, req *upkg.RegisterRequest) (res *upkg.RegisterResponse, err error)
}

type model struct {
}

func NewModel() Model {
	return &model{}
}

func (m *model) Register(c echo.Context, req *upkg.RegisterRequest) (res *upkg.RegisterResponse, err error) {
	// Generate Random String for Password
	req.Password = helper.RandStringBytes(4)

	stmt, err := db.SQLite.Conn.Prepare(`
		INSERT INTO "main"."users" (
			"name", "phone", "password", "role"
		) VALUES (?,?,?,?);
	`)

	if err != nil {
		c.Logger().Errorf("Error prepared query insert users :", err.Error())
		return nil, err
	}

	defer stmt.Close()

	_, err = stmt.Exec(req.Name, req.Phone, req.Password, req.Role)
	if err != nil {
		c.Logger().Errorf("Error executed query insert users :", err.Error())
		return nil, err
	}
	defer stmt.Close()

	return &upkg.RegisterResponse{
		Name:     req.Name,
		Phone:    req.Phone,
		Password: req.Password,
		Role:     req.Role,
	}, nil
}
