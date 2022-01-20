package controllers

import (
	"fmt"
	"project-e-commerces/constants"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

//================
//TOKEN ACCESS
//================
func CreateToken(userId int, role string) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	if role == "admin" {
		claims["admin"] = true
	} else {
		claims["admin"] = false
	}
	claims["userId"] = int(userId)
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(constants.JWT_SECRET_KEY))
}

func ExtractTokenUserId(c echo.Context) int {
	user := c.Get("user").(*jwt.Token)
	// fmt.Println(user)
	if user.Valid {
		claims := user.Claims.(jwt.MapClaims)
		userId := int(claims["userId"].(float64))
		fmt.Println(claims)
		return userId
	}
	return 0
}

// func checkAdmin(c echo.Context) string {
// 	user := c.Get("user").(*jwt.Token)
// 	// fmt.Println(user)
// 	if user.Valid {
// 		claims := user.Claims.(jwt.MapClaims)
// 		role := claims["admin"].(string)
// 		// fmt.Println(claims)

// 		return role
// 	}
// 	return "member"
// }

var IsAuthenticated = middleware.JWTWithConfig(middleware.JWTConfig{
	SigningKey: []byte("inirahasiayagaes"),
})

func IsAdmin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := c.Get("user").(*jwt.Token)
		claims := user.Claims.(jwt.MapClaims)
		admin := claims["admin"].(bool)
		if !admin {
			return echo.ErrUnauthorized
		}
		return next(c)
	}
}

// func IsAdmin(c echo.Context) error {
// 	role := ExtractTokenUserId(c)

// }
