package middleware

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"

	"github.com/golang-jwt/jwt/v4"
	"github.com/rindangramadhan/efishery-be-task/pkg/response"
)

var unlessPath = []string{
	"/api/health/status",
	"/api/users/register",
	"/api/users/login",
}

var unlessPrefix = []string{}

func JWT(h http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			// var err error
			var isUnlessPath, isUnlessPrefix bool

			for _, val := range unlessPrefix {
				regex := regexp.MustCompile(`^` + val)
				if regex.MatchString(r.URL.Path) {
					isUnlessPrefix = true
					break
				}
			}

			for _, v := range unlessPath {
				if r.URL.Path == v {
					isUnlessPath = !isUnlessPath
					break
				}
			}

			var session string

			authHeader := r.Header.Get("Authorization")
			session = strings.Replace(authHeader, "Bearer ", "", -1)

			// Unless Prefix
			if isUnlessPrefix && session == "" {
				h.ServeHTTP(w, r)
				return
			}

			// Unless Path
			if isUnlessPath && session == "" {
				h.ServeHTTP(w, r)
				return
			}

			if !strings.Contains(authHeader, "Bearer") {
				log.Println("Invalid token")
				WriteError(w, http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized))
				return
			}

			token, err := jwt.Parse(session, func(token *jwt.Token) (interface{}, error) {
				// ! Don't forget to validate the alg what you expect
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
				}

				return []byte(os.Getenv("SESSION_SECRET")), nil
			})

			if err != nil {
				log.Println("Error parse token :", err.Error())
				WriteError(w, http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized))
				return
			}

			if token == nil {
				log.Println("Error token empty :", err.Error())
				WriteError(w, http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized))
				return
			}

			_, ok := token.Claims.(jwt.MapClaims)
			if !ok || !token.Valid {
				log.Println("Error claims token :", err.Error())
				WriteError(w, http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized))
				return
			}

			h.ServeHTTP(w, r)
		},
	)
}

func WriteError(w http.ResponseWriter, code int, message string) (resp response.BodyTpl) {
	resp = response.BodyTpl{
		Status:  "error",
		Message: message,
	}

	result, _ := json.Marshal(resp)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(result)

	return
}
