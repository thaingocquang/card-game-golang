package util

import (
	"card-game-golang/config"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"time"
)

// JwtCustomClaims ...
type JwtCustomClaims struct {
	//ID string
	Data map[string]interface{}
	jwt.StandardClaims
}

// envVars ...
var envVars = config.GetEnv()

// GenerateUserToken ...
func GenerateUserToken(data map[string]interface{}) (string, error) {
	// claims ...
	claims := &JwtCustomClaims{
		data,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 120).Unix(),
		},
	}

	// generate token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// sign token
	st, err := token.SignedString([]byte(envVars.Jwt.SecretKey))

	// if err
	if err != nil {
		return "", err
	}

	return st, nil
}

// GetJWTPayload ...
func GetJWTPayload(c echo.Context) (map[string]interface{}, error) {
	// get jwt object from context
	user := c.Get("user").(*jwt.Token)

	claims := &JwtCustomClaims{}

	// ParseWithClaims
	_, err := jwt.ParseWithClaims(user.Raw, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.GetEnv().Jwt.SecretKey), nil
	})

	// if err
	if err != nil {
		return nil, err
	}

	return claims.Data, nil
}
