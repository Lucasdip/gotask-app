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