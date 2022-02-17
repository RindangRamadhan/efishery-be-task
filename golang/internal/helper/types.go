package helper

import (
	"math/rand"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type (
	Route struct {
		Method      string
		Pattern     string
		HandlerFunc echo.HandlerFunc
		Middleware  []echo.MiddlewareFunc
	}

	CustomBinder struct{}

	CustomValidator struct {
		Validator *validator.Validate
	}

	ValidationError struct {
		Field   string `json:"field"`
		Message string `json:"message"`
	}

	ValidationErrors []ValidationError
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func RandStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}
