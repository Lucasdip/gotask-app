package handlers

import (
	"os"
	"net/http"
	"time"
	"github.com/dgrijalva/jwt-go" // ou v5
	"github.com/gin-gonic/gin"
)

var secretKey = []byte(os.Getenv("JWT_SECRET"))

func getSecretKey() []byte {
    key := os.Getenv("JWT_SECRET")
    if key == "" {
        // Se não achar a chave, o app para aqui com um aviso claro
        panic("ERRO: A variável de ambiente JWT_SECRET não foi definida!")
    }
    return []byte(key)
}


func Login(c *gin.Context) {
	var credentials struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&credentials); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos"})
		return
	}

	// Login fixo para teste (Depois você pode buscar no banco)
	if credentials.Username == "admin" && credentials.Password == "1234" {
		// Criar o Token
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"user": credentials.Username,
			"exp":  time.Now().Add(time.Hour * 24).Unix(), // Expira em 24h
		})

		tokenString, err := token.SignedString(secretKey)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao gerar token"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"token": tokenString})
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Não autorizado"})
	}
}