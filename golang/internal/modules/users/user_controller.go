package users

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/rindangramadhan/efishery-be-task/internal/helper"
	"github.com/rindangramadhan/efishery-be-task/pkg/response"

	upkg "github.com/rindangramadhan/efishery-be-task/pkg/request/users"
)

type Controller interface {
	Register(ctx echo.Context) (err error)
	Login(ctx echo.Context) (err error)
	LoginCheck(ctx echo.Context) (err error)
}

type controller struct {
	m Model
}

func NewController() Controller {
	return &controller{
		m: NewModel(),
	}
}

// Register godoc
// @Tags         Register
// @Description  To create a new user
// @Accept       json
// @Produce      json
// @Param        payload  body      users.RegisterRequest  true  "payload"
// @Success      200      {object}  response.BodyTpl{status=string,message=string,values=response.Object}
// @Failure      400      {object}  response.BodyTpl{status=string,message=string,errors=response.Object}
// @Router       /auth/register [post]
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
		if strings.Contains(err.Error(), "UNIQUE constraint failed: users.username") {
			stackEntries := map[string]interface{}{
				"message": []interface{}{
					map[string]interface{}{
						"field":   "username",
						"message": "username already exist",
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

// Login godoc
// @Tags         Login
// @Description  To login user
// @Accept       json
// @Produce      json
// @Param        payload  body      users.LoginRequest  true  "payload"
// @Success      200      {object}  response.BodyTpl{status=string,message=string,values=response.Object}
// @Failure      400      {object}  response.BodyTpl{status=string,message=string,errors=response.Object}
// @Router       /auth/login [post]
func (c *controller) Login(ctx echo.Context) (err error) {
	// Bind & Validation Request
	req := new(upkg.LoginRequest)
	if err := ctx.Bind(req); err != nil {
		return response.WriteError(ctx, http.StatusBadRequest, http.StatusText(http.StatusBadRequest), err)
	}

	if err := ctx.Validate(req); err != nil {
		return response.WriteError(ctx, http.StatusBadRequest, http.StatusText(http.StatusBadRequest), err)
	}

	resp, err := c.m.Login(ctx, req)

	if err != nil {
		if strings.Contains(err.Error(), "sql: no rows in result set") {
			stackEntries := map[string]interface{}{
				"message": "Credentials not found",
			}

			return response.WriteError(ctx, http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized), stackEntries)
		} else {
			return response.WriteError(ctx, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), err)
		}
	}

	// Generate Token
	resp.Token, err = helper.GenerateToken(helper.TokenRequest{
		Name:      resp.Name,
		Phone:     resp.Phone,
		Role:      resp.Role,
		CreatedAt: resp.CreatedAt,
	})
	if err != nil {
		return response.WriteError(ctx, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), err)
	}

	return response.WriteSuccess(ctx, "Successfully Login", resp)
}

// LoginCheck godoc
// @Tags         Login Check
// @Description  To check login user
// @Accept       json
// @Produce      json
// @Security     ApiKeyAuth
// @Param        Authorization  header    string  true  "Insert your access token"  default(Bearer <Add access token here>)
// @Success      200            {object}  response.BodyTpl{status=string,message=string,values=response.Object}
// @Failure      401            {object}  response.BodyTpl{status=string,message=string,errors=response.Object}
// @Router       /auth/login-check [get]
func (c *controller) LoginCheck(ctx echo.Context) (err error) {
	authHeader := ctx.Request().Header.Get("Authorization")
	session := strings.Replace(authHeader, "Bearer ", "", -1)

	if !strings.Contains(authHeader, "Bearer") {
		log.Println("Invalid token")
		return response.WriteError(ctx, http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized), nil)
	}

	token, err := jwt.Parse(session, func(token *jwt.Token) (interface{}, error) {
		// ! Don't forget to validate the alg what you expect
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(os.Getenv("SESSION_SECRET")), nil
	})

	if err != nil {
		fmt.Println("Error parse token :", err.Error())
		return response.WriteError(ctx, http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized), nil)
	}

	if token == nil {
		fmt.Println("Error token empty :", err.Error())
		return response.WriteError(ctx, http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized), nil)
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		fmt.Println("Error claims token :", err.Error())
		return response.WriteError(ctx, http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized), nil)
	}

	return response.WriteSuccess(ctx, "Successfully Login Check", claims["sub"])
}
