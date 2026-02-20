package handlers

import (
	"os"
	"strings"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(401, gin.H{"error": "Caminho proibido para quem não é Dragonborn"})
			return
		}

		// Importante: O JS envia "Bearer <token>", temos que separar!
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.AbortWithStatusJSON(401, gin.H{"error": "Formato de pergaminho inválido"})
			return
		}

		tokenString := parts[1]

		// Valida o token usando a mesma Secret que você usou no Login
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_SECRET")), nil
		})

		if err != nil || !token.Valid {
			c.AbortWithStatusJSON(401, gin.H{"error": "Este pergaminho expirou ou é falso"})
			return
		}

		c.Next()
	}
}