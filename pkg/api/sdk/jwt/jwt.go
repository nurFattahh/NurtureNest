package crypto

import (
	"NurtureNest/pkg/api/sdk/response"
	"NurtureNest/pkg/config"
	"NurtureNest/pkg/domain"
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func GenerateToken(payload domain.Users) (string, error) {
	config, _ := config.LoadConfig()
	expStr := config.JWT_EXP
	var exp time.Duration
	exp, err := time.ParseDuration(expStr)
	if expStr == "" || err != nil {
		exp = time.Hour * 1
	}

	temporaryJwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":   payload.ID,
		"role": payload.Role,
		"exp":  time.Now().Add(exp).Unix(),
	})

	tokenJwt, err := temporaryJwtToken.SignedString([]byte(config.JWT_SECRET_KEY))
	if err != nil {
		return "", err
	}
	return tokenJwt, nil
}

func ClaimToken(c *gin.Context) (IDclaims float64, role string) {
	result, exist := c.Get("user")
	if !exist {
		response.FailOrError(c, http.StatusUnprocessableEntity, "no user key found", errors.New("Error"))
		return
	}
	claims, ok := result.(jwt.MapClaims)
	if !ok {
		response.FailOrError(c, http.StatusUnprocessableEntity, "error parsing ", errors.New("Error"))
		return
	}

	userIDc := claims["id"]
	userIDf, ok := userIDc.(float64)
	if !ok {
		response.FailOrError(c, http.StatusUnprocessableEntity, "error get id", errors.New("Error"))
		return
	}

	result, exist = c.Get("user")
	if !exist {
		response.FailOrError(c, http.StatusUnprocessableEntity, "no user key found", errors.New("Error"))
		return
	}
	claims, ok = result.(jwt.MapClaims)
	if !ok {
		response.FailOrError(c, http.StatusUnprocessableEntity, "error parsing ", errors.New("Error"))
		return
	}

	userRole := claims["role"]
	userRolef, ok := userRole.(string)
	if !ok {
		response.FailOrError(c, http.StatusUnprocessableEntity, "error get role", errors.New("Error"))
		return
	}

	return userIDf, userRolef
}
