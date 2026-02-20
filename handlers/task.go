package handlers

import (
    "net/http"
    "gotask-app/models"
    "github.com/gin-gonic/gin"
)

// Listar todas as tasks do banco
func GetTasks(c *gin.Context) {
	var tasks []models.Task
	models.DB.Find(&tasks) // SELECT * FROM tasks
	c.JSON(http.StatusOK, tasks)
}

// Criar uma nova task no banco
func CreateTask(c *gin.Context) {
	var newTask models.Task
	if err := c.ShouldBindJSON(&newTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := models.DB.Create(&newTask) // INSERT INTO tasks...
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao salvar"})
		return
	}

	c.JSON(http.StatusCreated, newTask)
}

// Marcar Task como feita ou mudar o título (Update)
func UpdateTask(c *gin.Context) {
	id := c.Param("id") // Pega o ID da URL, ex: /tasks/5
	var task models.Task

	// 1. Tentar encontrar a task no banco pelo ID
	if err := models.DB.First(&task, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task não encontrada"})
		return
	}

	// 2. Mapear o que veio no JSON para a task encontrada
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 3. Salvar as mudanças no banco
	models.DB.Save(&task)
	c.JSON(http.StatusOK, task)
}

// Apagar uma Task do banco (Delete)
func DeleteTask(c *gin.Context) {
	id := c.Param("id")
	var task models.Task

	// Tentar deletar diretamente pelo ID
	result := models.DB.Delete(&task, id)

	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task não encontrada ou já deletada"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Task removida com sucesso!"})
}