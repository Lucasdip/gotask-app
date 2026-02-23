package handlers

import (
    "net/http"
    "gotask-app/models"
    "github.com/gin-gonic/gin"
)

// Listar todas as transações do usuário
func GetTransactions(c *gin.Context) {
    userID, _ := c.Get("userID")
    var transactions []models.Transaction

    models.DB.Where("user_id = ?", userID).Find(&transactions)
    c.JSON(http.StatusOK, transactions)
}

// Criar nova transação (Gasto ou Ganho)
func CreateTransaction(c *gin.Context) {
    userID, _ := c.Get("userID")
    var transaction models.Transaction

    if err := c.ShouldBindJSON(&transaction); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    transaction.UserID = userID.(uint)
    models.DB.Create(&transaction)
    c.JSON(http.StatusOK, transaction)
}

// Deletar transação
func DeleteTransaction(c *gin.Context) {
    id := c.Param("id")
    userID, _ := c.Get("userID")
    models.DB.Where("id = ? AND user_id = ?", id, userID).Delete(&models.Transaction{})
    c.JSON(http.StatusOK, gin.H{"message": "Transação removida"})
}