package middleware

import (
	"NurtureNest/pkg/api/sdk/response"
	"NurtureNest/pkg/config"
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func JwtMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		authorization := c.Request.Header.Get("Authorization")
		if !strings.HasPrefix(authorization, "Bearer ") {
			c.Abort()
			msg := "wrong header value"
			response.FailOrError(c, http.StatusForbidden, msg, errors.New(msg))
			return
		}

		config, _ := config.LoadConfig()
		tokenJwt := authorization[7:]
		validateJwt, err := jwt.Parse(tokenJwt, func(t *jwt.Token) (interface{}, error) {
			return []byte(config.JWT_SECRET_KEY), nil
		})

		if err != nil {
			c.Abort()
			response.FailOrError(c, http.StatusForbidden, err.Error(), err)
			return
		}

		jwtFix, ok := validateJwt.Claims.(jwt.MapClaims)
		if !ok {
			c.Abort()
			response.FailOrError(c, http.StatusForbidden, "data token jwt tidak valid", nil)
			return
		}

		if jwtFix.Valid() != nil {
			c.Abort()
			response.FailOrError(c, http.StatusForbidden, jwtFix.Valid().Error(), jwtFix.Valid())
			return
		} else {

			c.Set("user", jwtFix)
			c.Next()
		}
	}
}
