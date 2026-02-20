package middlewares

import (
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "authorization header is required",
			})
			c.Abort()
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "invalid authorization format",
			})
			c.Abort()
			return
		}

		tokenString := parts[1]
		secret := os.Getenv("JWT_SECRET")
		if secret == "" {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "server configuration error",
			})
			c.Abort()
			return
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrSignatureInvalid
			}
			return []byte(secret), nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "invalid or expired token",
			})
			c.Abort()
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "invalid token claims",
			})
			c.Abort()
			return
		}

		// 🔑 Extract important claims
		userID, okID := claims["user_id"].(float64)
		userEmail, okEmail := claims["username"].(string)
		role, okRole := claims["role"].(string)

		if !okID || !okEmail || !okRole {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "invalid token payload",
			})
			c.Abort()
			return
		}

		// 📌 Simpan ke context
		c.Set("user_id", int(userID))
		c.Set("user_email", userEmail)
		c.Set("role", role)

		// Optional
		if email, ok := claims["email"].(string); ok {
			c.Set("email", email)
		}

		c.Next()
	}
}

func AdminOnly() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Println("check the keys : ", c.Keys)
		if c.GetString("role") != "admin" {
			c.JSON(http.StatusForbidden, gin.H{
				"message": "admin access only",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
