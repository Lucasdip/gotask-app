package handlers

import (
	"github.com/gin-gonic/gin"
	"gotask-app/models"
	"net/http"
)

// Listar todas as tasks do banco
func GetTasks(c *gin.Context) {
	userID, _ := c.Get("userID")

	var user models.User
	models.DB.First(&user, userID)

	var tasks []models.Task
	models.DB.Where("user_id = ?", userID).Find(&tasks)

	// Aqui usamos o calculateRank que acabamos de criar
	rank := calculateRank(user.Level)

	c.JSON(http.StatusOK, gin.H{
		"tasks": tasks,
		"user": gin.H{
			"level": user.Level,
			"xp":    user.XP,
			"rank":  rank,
		},
	})
}

// Criar uma nova task no banco
func CreateTask(c *gin.Context) {
	// 1. Pega o ID do usuário logado através do middleware
	userID, _ := c.Get("userID")

	var newTask models.Task
	if err := c.ShouldBindJSON(&newTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 2. VINCULA A TASK AO USUÁRIO (O que estava faltando)
	newTask.UserID = userID.(uint)

	result := models.DB.Create(&newTask)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao salvar no diário de quests"})
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
    userID, _ := c.Get("userID") // Pega o ID do usuário atual

    var task models.Task
    // Só deleta se a task pertencer ao usuário logado
    result := models.DB.Where("id = ? AND user_id = ?", id, userID).Delete(&task)

    if result.RowsAffected == 0 {
        c.JSON(http.StatusNotFound, gin.H{"error": "Quest não encontrada ou você não tem permissão"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Quest abandonada com sucesso!"})
}
func ToggleTask(c *gin.Context) {
	id := c.Param("id")
	userID, _ := c.Get("userID")

	var task models.Task
	if err := models.DB.Where("id = ? AND user_id = ?", id, userID).First(&task).Error; err != nil {
		c.JSON(404, gin.H{"error": "Quest não encontrada"})
		return
	}

	// Só ganha XP se a task estiver sendo marcada como "concluída" agora
	wasDone := task.Done
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(400, gin.H{"error": "Dados inválidos"})
		return
	}

	if !wasDone && task.Done {
		// Lógica de Ganho de XP
		var user models.User
		models.DB.First(&user, userID)

		user.XP += 20 // 20 de XP por missão

		// Lógica de Level Up (a cada 100 XP sobe um nível)
		if user.XP >= 100 {
			user.Level += 1
			user.XP = 0 // Reseta o XP para o próximo nível
		}

		models.DB.Save(&user)
	}

	models.DB.Save(&task)
	c.JSON(200, task)
}

// Função para definir a patente com base no nível (Regra de 10 em 10)
func calculateRank(level int) string {
	if level <= 10 {
		return "Prisioneiro de Helgen"
	} else if level <= 20 {
		return "Aprendiz de Winterhold"
	} else if level <= 30 {
		return "Thane de Whiterun"
	} else if level <= 40 {
		return "Líder da Guilda"
	}
	return "Dovahkiin"
}
