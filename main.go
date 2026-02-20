package main

import (
	"fmt"
    "gotask-app/handlers"
	"gotask-app/models"
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


	// Criamos um grupo protegido
	authorized := r.Group("/tasks")
	authorized.Use(handlers.AuthMiddleware()) // <--- O Pedágio está aqui!
	{
		authorized.GET("/", handlers.GetTasks)
		authorized.POST("/", handlers.CreateTask)
		authorized.PUT("/:id", handlers.UpdateTask)
		authorized.DELETE("/:id", handlers.DeleteTask)
	}

	r.Run(":8080")

}