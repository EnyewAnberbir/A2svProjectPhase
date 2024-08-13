package infrastructure

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/dgrijalva/jwt-go"
)

type JWTAuthMiddleware struct {
	jwtService JWTService
}

func NewJWTAuthMiddleware(jwtService JWTService) *JWTAuthMiddleware {
	return &JWTAuthMiddleware{jwtService: jwtService}
}

func (j *JWTAuthMiddleware) AuthenticateJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
			c.Abort()
			return
		}

		if !strings.HasPrefix(authHeader, "Bearer ") {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization header format"})
			c.Abort()
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		token, err := j.jwtService.ValidateToken(tokenString)
		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
			c.Abort()
			return
		}

		c.Set("userID", claims["sub"].(string))
		c.Set("role", claims["role"].(string))
		c.Next()
	}
}

func (j *JWTAuthMiddleware) AuthorizeAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		role, exists := c.Get("role")
		if !exists || role != "admin" {
			c.JSON(http.StatusForbidden, gin.H{"error": "Admin access required"})
			c.Abort()
			return
		}
		c.Next()
	}
}
