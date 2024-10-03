package middlewares

import (
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/maulanadityaa/bank-merchant-api/models/dto/response"
	"github.com/maulanadityaa/bank-merchant-api/repositories"
	"github.com/maulanadityaa/bank-merchant-api/repositories/impl"
	"github.com/maulanadityaa/bank-merchant-api/utils"
)

var (
	jwtSecret                                            = os.Getenv("JWT_SECRET")
	jwtSignatureKey     []byte                           = []byte(jwtSecret)
	blacklistRepository repositories.BlacklistRepository = impl.NewBlacklistRepository()
)

func AuthWithRole(role []string) gin.HandlerFunc {
	return func(c *gin.Context) {
		jwtClaims := utils.GetJWTClaims(c)

		userRole := jwtClaims["role"].(string)

		for _, r := range role {
			if userRole == r {
				c.Next()
				return
			}
		}

		response.NewResponseForbidden(c, "You don't have permission to access this endpoint")
		c.Abort()
	}
}

func ValidateJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			response.NewResponseUnauthorized(c, "Token is required")
			c.Abort()
			return
		}

		parts := strings.Split(tokenString, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			response.NewResponseUnauthorized(c, "Invalid token")
			c.Abort()
			return
		}

		tokenString = parts[1]

		isBlacklist, _ := blacklistRepository.IsBlacklist(tokenString)
		if isBlacklist {
			response.NewResponseUnauthorized(c, "Token is blacklisted")
			c.Abort()
			return
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrSignatureInvalid
			}

			return jwtSignatureKey, nil
		})

		if err != nil {
			response.NewResponseUnauthorized(c, err.Error())
			c.Abort()
			return
		}

		if !token.Valid {
			response.NewResponseUnauthorized(c, "Invalid token")
			c.Abort()
			return
		}

		c.Next()
	}
}
