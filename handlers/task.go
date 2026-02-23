package handlers

import (
	"gotask-app/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Listar tarefas filtradas pelo dono do Token + Dados do Usuário (XP/Level)
func GetTasks(c *gin.Context) {
    // 1. Pegar o ID do contexto (salvo pelo AuthMiddleware)
    val, exists := c.Get("userID")
    if !exists {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "ID do herói não encontrado"})
        return
    }

    // Converter para uint com segurança
    userID := val.(uint)

    var tasks []models.Task
    var user models.User

    // 2. Buscar tarefas do usuário
    if err := models.DB.Where("user_id = ?", userID).Find(&tasks).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro nas quests"})
        return
    }

    // 3. Buscar dados do usuário (XP, Level, etc)
    if err := models.DB.First(&user, userID).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Herói não encontrado"})
        return
    }

    // 4. Retornar exatamente o que o JS espera
    c.JSON(http.StatusOK, gin.H{
        "tasks": tasks,
        "user":  user,
    })
}
// Criar uma nova task vinculada ao herói logado
func CreateTask(c *gin.Context) {
	userID, _ := c.Get("userID")
	var newTask models.Task

	if err := c.ShouldBindJSON(&newTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// VINCULA o ID do usuário à nova tarefa
	newTask.UserID = userID.(uint)

	if err := models.DB.Create(&newTask).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao salvar quest"})
		return
	}

	c.JSON(http.StatusCreated, newTask)
}

// Lógica de Toggle (Concluir Missão e Ganhar XP)
func ToggleTask(c *gin.Context) {
	id := c.Param("id")
	userID, _ := c.Get("userID")
	var task models.Task
	var user models.User

	// 1. Acha a tarefa garantindo que ela é do usuário logado
	if err := models.DB.Where("id = ? AND user_id = ?", id, userID).First(&task).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Quest não encontrada"})
		return
	}

	// 2. Salva o status antigo para saber se deve dar XP
	wasDone := task.Done

	// 3. Atualiza com o novo status vindo do front
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos"})
		return
	}

	// 4. Se a quest era falsa e virou verdadeira: LEVEL UP LOGIC!
	if !wasDone && task.Done {
		models.DB.First(&user, userID)
		user.XP += 20

		// Sobe de nível se chegar em 100
		for user.XP >= 100 {
			user.Level++
			user.XP -= 100
			
			// Atualiza o Rank baseado no novo nível
			if user.Level >= 5 { user.Rank = "Thane de Whiterun" }
			if user.Level >= 10 { user.Rank = "Líder da Guilda" }
		}
		models.DB.Save(&user)
	}

	models.DB.Save(&task)
	c.JSON(http.StatusOK, gin.H{"task": task, "user": user})
}

// Apagar uma Task
func DeleteTask(c *gin.Context) {
	id := c.Param("id")
	userID, _ := c.Get("userID")

	// Só deleta se a task for do usuário
	result := models.DB.Where("id = ? AND user_id = ?", id, userID).Delete(&models.Task{})

	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Quest não encontrada"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Quest abandonada!"})
}