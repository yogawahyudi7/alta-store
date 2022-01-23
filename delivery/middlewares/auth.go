package middlewares

import (
	"errors"
	"net/http"
	"strings"
	"time"

	// jwtv3 "github.com/dgrijalva/jwt-go/v4"
	jwtv3 "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

type Auth interface {
	GenerateToken(userID int) (string, error)
	// ExtractTokenUserID(e echo.Context) int
}

type jwtService struct {
}

var SecretKey = "secret123"

func NewAuth() *jwtService {
	return &jwtService{}
}

func (a *jwtService) GenerateToken(userID int, role string) (string, error) {
	claim := jwtv3.MapClaims{}
	claim["user_id"] = userID
	claim["exp"] = time.Now().Add(time.Hour * 1).Unix()

	if role == "admin" {
		claim["admin"] = true
	} else {
		claim["admin"] = false
	}

	token := jwtv3.NewWithClaims(jwtv3.SigningMethodHS256, claim)

	signedToken, err := token.SignedString([]byte(SecretKey))
	if err != nil {
		return signedToken, err
	}

	return signedToken, nil
}

func (a *jwtService) ExtractTokenUserID(c echo.Context) int {
	authHeader := c.Request().Header.Get(echo.HeaderAuthorization)

	if !strings.Contains(authHeader, "Bearer") {
		return 0
	}

	tokenString := ""
	arrayToken := strings.Split(authHeader, " ")
	if len(arrayToken) == 2 {
		tokenString = arrayToken[1]
	}

	token, _ := jwtv3.Parse(tokenString, func(token *jwtv3.Token) (interface{}, error) {
		_, ok := token.Method.(*jwtv3.SigningMethodHMAC)

		if !ok {
			return nil, errors.New("invalid token")
		}

		return []byte(SecretKey), nil
	})

	claim, ok := token.Claims.(jwtv3.MapClaims)

	if !ok || !token.Valid {
		return 0
	}

	userID := int(claim["user_id"].(float64))

	return userID

}

func (a *jwtService) IsAdmin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authHeader := c.Request().Header.Get(echo.HeaderAuthorization)

		tokenString := ""
		arrayToken := strings.Split(authHeader, " ")
		if len(arrayToken) == 2 {
			tokenString = arrayToken[1]
		}

		token, _ := jwtv3.Parse(tokenString, func(token *jwtv3.Token) (interface{}, error) {
			_, ok := token.Method.(*jwtv3.SigningMethodHMAC)

			if !ok {
				return nil, errors.New("invalid token")
			}

			return []byte(SecretKey), nil
		})

		claim, _ := token.Claims.(jwtv3.MapClaims)

		admin := claim["admin"].(bool)
		if !admin {
			return c.JSON(http.StatusUnauthorized, map[string]interface{}{
				"code":    http.StatusUnauthorized,
				"message": "not admin",
			})
		}

		return next(c)
	}
}
