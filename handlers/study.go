package handlers

import (
	"gotask-app/models"
	"net/http"
	"github.com/gin-gonic/gin"
)

// Função auxiliar interna para evitar repetição de código
func getUserID(c *gin.Context) (uint, bool) {
	uid, exists := c.Get("userID")
	if !exists {
		return 0, false
	}
	switch v := uid.(type) {
	case uint:
		return v, true
	case float64:
		return uint(v), true
	default:
		return 0, false
	}
}

func GetStudies(c *gin.Context) {
    userID, ok := getUserID(c)
    if !ok {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Usuário não autorizado"})
        return
    }

    // Inicializando como slice vazia em vez de nil
    studies := []models.Study{} 
    models.DB.Where("user_id = ?", userID).Find(&studies)
    c.JSON(http.StatusOK, studies)
}

func CreateStudy(c *gin.Context) {
	var input models.Study
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos"})
		return
	}

	userID, ok := getUserID(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Usuário não autorizado"})
		return
	}

	input.UserID = userID
	input.Status = "Ativo"

	if err := models.DB.Create(&input).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao criar estudo"})
		return
	}
	c.JSON(http.StatusCreated, input)
}

func DeleteStudy(c *gin.Context) {
	id := c.Param("id")
	userID, ok := getUserID(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Usuário não autorizado"})
		return
	}

	// Deleta apenas se o estudo pertencer ao usuário logado (Segurança!)
	result := models.DB.Where("id = ? AND user_id = ?", id, userID).Delete(&models.Study{})
	
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao deletar"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Estudo removido do grimório"})
}