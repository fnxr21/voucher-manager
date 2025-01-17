package middleware

import (
	// "log"
	jwtToken "github.com/fnxr21/item-list/pkg/jwt"
	resultType "github.com/fnxr21/item-list/pkg/type"
	"log"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

// Declare Result struct here ...
type Result struct {
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

// Create Auth function here ...
func Auth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Request().Header.Get("Authorization")

		if token == "" {
			return c.JSON(http.StatusUnauthorized, resultType.ErrorResult{Status: http.StatusBadRequest, Message: "unauthorized-token"})
		}

		token = strings.Split(token, " ")[1]
		claims, err := jwtToken.DecodeToken(token)

		if err != nil {
			return c.JSON(http.StatusUnauthorized, resultType.ErrorResult{Status: http.StatusUnauthorized, Message: "unauthorized"})
		}
			log.Println("user", claims)
			c.Set("userLogin", claims)
			return next(c)
		

	}
}

func GetUserIdFromContext(c echo.Context) int {
	if c == nil {
		return -1
	}

	// Check user login
	if userLogin, ok := c.Get("userLogin").(jwt.MapClaims); ok {
		if id, ok := userLogin["id"].(float64); ok {
			return int(id)
		}
	}

	return -1 // Return -1 if no valid login found
}
