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
	reqPassword := helper.RandStringBytes(4)

	stmt, err := db.SQLite.Conn.Prepare(`
		INSERT INTO "main"."users" (
			"username", "name", "phone", "password", "role", "created_at"
		) VALUES (?,?,?,?,?,?);
	`)

	if err != nil {
		c.Logger().Errorf("Error prepared query insert users :", err.Error())
		return nil, err
	}

	defer stmt.Close()

	_, err = stmt.Exec(req.Username, req.Name, req.Phone, reqPassword, req.Role, currTime)
	if err != nil {
		c.Logger().Errorf("Error executed query insert users :", err.Error())
		return nil, err
	}
	defer stmt.Close()

	return &upkg.RegisterResponse{
		Username: req.Username,
		Name:     req.Name,
		Phone:    req.Phone,
		Password: reqPassword,
		Role:     req.Role,
	}, nil
}

func (m *model) Login(c echo.Context, req *upkg.LoginRequest) (res *upkg.LoginResponse, err error) {
	var row struct {
		Name      string `json:"name"`
		Role      string `json:"role"`
		CreatedAt string `json:"created_at"`
	}

	stmt, err := db.SQLite.Conn.Prepare(`
		SELECT name, role, created_at FROM "main"."users" 
			WHERE phone = ? AND password = ? 
	`)

	if err != nil {
		c.Logger().Errorf("Error prepared query select users :", err.Error())
		return nil, err
	}

	defer stmt.Close()

	err = stmt.QueryRow(req.Phone, req.Password).Scan(&row.Name, &row.Role, &row.CreatedAt)
	if err != nil {
		c.Logger().Errorf("Error executed query select users :", err.Error())
		return nil, err
	}
	defer stmt.Close()

	return &upkg.LoginResponse{
		Name:      row.Name,
		Phone:     req.Phone,
		Role:      row.Role,
		CreatedAt: row.CreatedAt,
	}, nil
}
