package middleware

import (
	"fmt"
	"net/http"
	"service"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func AuthorizeJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		const BEARER_SCHEMA = "Bearer"
		// Get the Authorization information in the header
		if authHeader := c.GetHeader("Authorization"); len(authHeader) > 0 {
			tokenString := strings.TrimSpace(authHeader[len(BEARER_SCHEMA):])

			// Token validation
			token, err := service.JWTAuthService().ValidateToken(tokenString)

			if token.Valid {
				claims := token.Claims.(jwt.MapClaims)
				fmt.Println(claims)
			} else {
				fmt.Println(err)
				c.AbortWithStatus(http.StatusUnauthorized)
			}
		} else {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

	}
}
