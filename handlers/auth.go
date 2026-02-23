package handlers

import (
	"os"
	"net/http"
	"gotask-app/models"
	"time"
	"github.com/dgrijalva/jwt-go" // ou v5
	"github.com/gin-gonic/gin"
)

var secretKey = []byte(os.Getenv("JWT_SECRET"))

// Função auxiliar para garantir que a chave nunca esteja vazia
func getSecretKey() []byte {
	key := os.Getenv("JWT_SECRET")
	if key == "" {
		// No Railway, se você esquecer de por a variável, o log vai te avisar
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

	// --- LÓGICA COM BANCO DE DADOS ---
	var user models.User
	// Busca no banco de dados o usuário enviado no JSON
	result := models.DB.Where("username = ?", credentials.Username).First(&user)

	// Verifica se o usuário existe E se a senha coincide
	if result.Error != nil || user.Password != credentials.Password {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Nome de herói ou grito de guerra incorretos"})
		return
	}

	// --- GERAÇÃO DO TOKEN ---
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "user_id":  user.ID,       // <--- ADICIONE ESTA LINHA (Fundamental!)
        "username": user.Username,
        "exp":      time.Now().Add(time.Hour * 24).Unix(),
    })

	// Usamos a função getSecretKey() aqui para garantir que pegamos o valor atualizado
	tokenString, err := token.SignedString(getSecretKey())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao gerar pergaminho de acesso"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}