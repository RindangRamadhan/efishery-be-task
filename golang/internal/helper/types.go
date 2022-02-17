package helper

import (
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
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

	TokenRequest struct {
		Name      string `json:"name"`
		Phone     string `json:"phone"`
		Role      string `json:"role"`
		CreatedAt string `json:"created_at"`
	}
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func RandStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func GenerateToken(req TokenRequest) (session string, err error) {
	sesExp, _ := strconv.Atoi(os.Getenv("SESSION_EXP"))
	expirationTime := time.Now().Add(time.Duration(sesExp) * time.Minute)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": req,
		"nbf": time.Now().Unix(),
		"exp": expirationTime.Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenStr, err := token.SignedString([]byte(os.Getenv("SESSION_SECRET")))
	if err != nil {
		return
	}

	session = tokenStr
	return
}
