package main

import (
	"fmt"
	"gotask-app/handlers"
	"gotask-app/models"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	godotenv.Load()
	
	err := godotenv.Load()
    if err != nil {
        fmt.Println("Aviso: Arquivo .env não encontrado, usando variáveis do sistema.")
    }
	
	models.InitDB()
	// No database.go ou main.go
	models.DB.AutoMigrate(&models.User{}, &models.Task{}, &models.Transaction{})

    r := gin.Default()
	r.StaticFile("/", "./index.html")
	
	r.Use(func(c *gin.Context) {
    c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
    c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
    c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
    
    if c.Request.Method == "OPTIONS" {
        c.AbortWithStatus(204)
        return
    }
    c.Next()
})


    r.POST("/login", handlers.Login)

    authorized := r.Group("/api")
    authorized.Use(handlers.AuthMiddleware())
    {
        // Tasks
        authorized.GET("/tasks", handlers.GetTasks)
        authorized.POST("/tasks", handlers.CreateTask)
        authorized.PUT("/tasks/:id", handlers.ToggleTask)
        authorized.DELETE("/tasks/:id", handlers.DeleteTask)
        
        // Finance
        authorized.GET("/finance", handlers.GetTransactions)
        authorized.POST("/finance", handlers.CreateTransaction)
        authorized.DELETE("/finance/:id", handlers.DeleteTransaction)
        
        // Studies
        authorized.GET("/studies", handlers.GetStudies)
        authorized.POST("/studies", handlers.CreateStudy)
        authorized.DELETE("/studies/:id", handlers.DeleteStudy)
    }

    port := os.Getenv("PORT")
    if port == "" { port = "8080" }

    fmt.Println("Servidor rodando na porta", port)
    r.Run(":" + port)
}