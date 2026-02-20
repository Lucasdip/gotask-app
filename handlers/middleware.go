package handlers

import (
	"net/http"
	"os"
	"strings"
    "gotask-app/models" // Ajuste para o caminho do seu projeto
    "github.com/dgrijalva/jwt-go"
    "github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        authHeader := c.GetHeader("Authorization")
        if authHeader == "" {
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Token não fornecido"})
            return
        }

        // Separa "Bearer <token>"
        parts := strings.Split(authHeader, " ")
        if len(parts) != 2 || parts[0] != "Bearer" {
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Formato de token inválido"})
            return
        }

        tokenString := parts[1]
        token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
            return []byte(os.Getenv("JWT_SECRET")), nil
        })

        if err != nil || !token.Valid {
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Token inválido ou expirado"})
            return
        }

        claims := token.Claims.(jwt.MapClaims)
        username := claims["user"].(string)

        // Busca o usuário completo no banco para pegar o ID e o nível atual
        var user models.User
        if err := models.DB.Where("username = ?", username).First(&user).Error; err != nil {
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Usuário não encontrado"})
            return
        }

        // Salva o ID no contexto para ser usado nos Handlers
        c.Set("userID", user.ID)
        c.Next()
    }
}