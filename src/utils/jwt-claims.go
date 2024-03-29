package utils

import (
	"merchandise-circulation-api/src/configs"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

type JwtCustomClaims struct {
	jwt.StandardClaims
}

func GenerateJwt(userId string) (token string, err error) {
	claims := JwtCustomClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
			Id:        userId,
		},
	}

	config, _ := configs.LoadServerConfig(".")
	rawToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err = rawToken.SignedString([]byte(config.JWTsecret))
	return
}

func SetJwtCookie(c echo.Context, token string) {
	authCookie := http.Cookie{
		Name:     "token",
		Value:    token,
		Expires:  time.Now().Add(24 * time.Hour),
		Path:     "/",
		HttpOnly: true,
	}
	c.SetCookie(&authCookie)
}

func GetJwtTokenFromRequest(c echo.Context) string {
	authHeader := c.Request().Header[echo.HeaderAuthorization][0]
	authToken := strings.Split(authHeader, "Bearer ")[1]
	return authToken
}

func ExtractClaims(tokenStr string) (*JwtCustomClaims, error) {
	config, _ := configs.LoadServerConfig(".")
	hmacSecretString := config.JWTsecret
	var claims *JwtCustomClaims

	token, err := jwt.ParseWithClaims(tokenStr, &JwtCustomClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(hmacSecretString), nil
	})

	if token != nil {
		if claims, ok := token.Claims.(*JwtCustomClaims); ok && token.Valid {
			return claims, nil
		}
	}

	return claims, err
}
