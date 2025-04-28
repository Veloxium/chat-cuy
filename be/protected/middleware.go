package protected

import (
	"github.com/Gylmynnn/websocket-sesat/utils"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"strings"
	"time"
)

var secretKey = utils.LoadENV("JWTSECRETKEY")

func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		var getToken string

		// get authorization from header
		authHeader := c.GetHeader("Authorization")
		if authHeader != "" {
			tokenFormat := strings.Split(authHeader, " ")
			if len(tokenFormat) == 2 && tokenFormat[0] == "Bearer" {
				getToken = tokenFormat[1]
			}
		}

		if getToken == "" {
			getTokenFromCookie, err := c.Cookie("jwt")
			if err == nil {
				getToken = getTokenFromCookie
			}
		}

		if getToken == "" {
			c.JSON(http.StatusUnauthorized, utils.ResFormatter{
				Success:    false,
				StatusCode: http.StatusUnauthorized,
				Message:    "unauthorized required token",
				Data:       nil,
			})
			c.Abort()
			return
		}

		token, err := jwt.ParseWithClaims(getToken, &utils.MyJWTClaims{}, func(t *jwt.Token) (any, error) {
			return []byte(secretKey), nil
		})
		if err != nil {
			c.JSON(http.StatusUnauthorized, utils.ResFormatter{
				Success:    false,
				StatusCode: http.StatusUnauthorized,
				Message:    "token invalid" + err.Error(),
				Data:       nil,
			})
			c.Abort()
			return
		}
		if claims, ok := token.Claims.(*utils.MyJWTClaims); ok && token.Valid {
			if claims.ExpiresAt.Before(time.Now()) {
				c.JSON(http.StatusUnauthorized, utils.ResFormatter{
					Success:    false,
					StatusCode: http.StatusUnauthorized,
					Message:    "token expired",
					Data:       nil,
				})
				c.Abort()
				return
			} else {
				c.Set("id", claims.ID)
				c.Set("username", claims.Username)
			}
		} else {
			c.JSON(http.StatusUnauthorized, utils.ResFormatter{
				Success:    false,
				StatusCode: http.StatusUnauthorized,
				Message:    "token invalid 2",
				Data:       nil,
			})
			c.Abort()
			return
		}
		c.Next()
	}

}
