package users

import (
	"time"

	"github.com/labstack/echo/v4"

	db "github.com/rindangramadhan/efishery-be-task/internal/db/sqlite"
	"github.com/rindangramadhan/efishery-be-task/internal/helper"
	upkg "github.com/rindangramadhan/efishery-be-task/pkg/request/users"
)

type Model interface {
	Register(c echo.Context, req *upkg.RegisterRequest) (res *upkg.RegisterResponse, err error)
	Login(c echo.Context, req *upkg.LoginRequest) (res *upkg.LoginResponse, err error)
}

type model struct {
}

func NewModel() Model {
	return &model{}
}

func (m *model) Register(c echo.Context, req *upkg.RegisterRequest) (res *upkg.RegisterResponse, err error) {
	var location, _ = time.LoadLocation("Asia/Jakarta")
	var currTime = time.Now().In(location).Format("2006-01-02T15:04:05-0700")

	// Generate Random String for Password
	req.Password = helper.RandStringBytes(4)

	stmt, err := db.SQLite.Conn.Prepare(`
		INSERT INTO "main"."users" (
			"name", "phone", "password", "role", "created_at"
		) VALUES (?,?,?,?,?);
	`)

	if err != nil {
		c.Logger().Errorf("Error prepared query insert users :", err.Error())
		return nil, err
	}

	defer stmt.Close()

	_, err = stmt.Exec(req.Name, req.Phone, req.Password, req.Role, currTime)
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

func (m *model) Login(c echo.Context, req *upkg.LoginRequest) (res *upkg.LoginResponse, err error) {
	var row struct {
		Phone     string `json:"phone"`
		Role      string `json:"role"`
		CreatedAt string `json:"created_at"`
	}

	stmt, err := db.SQLite.Conn.Prepare(`
		SELECT phone, role, created_at FROM "main"."users" 
			WHERE name = ? AND password = ? 
	`)

	if err != nil {
		c.Logger().Errorf("Error prepared query select users :", err.Error())
		return nil, err
	}

	defer stmt.Close()

	err = stmt.QueryRow(req.Name, req.Password).Scan(&row.Phone, &row.Role, &row.CreatedAt)
	if err != nil {
		c.Logger().Errorf("Error executed query select users :", err.Error())
		return nil, err
	}
	defer stmt.Close()

	return &upkg.LoginResponse{
		Name:      req.Name,
		Phone:     row.Phone,
		Password:  req.Password,
		Role:      row.Role,
		CreatedAt: row.CreatedAt,
	}, nil
}
