package models

import (
	"os"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)


var DB *gorm.DB

func InitDB() {

	dsn := os.Getenv("DATABASE_URL")

	if dsn == ""{
		panic("A variável DATABASE não foi encontrada na .env")
	}

	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Printf("Erro ao conectar no Neon: %v\n", err)
		panic(err)
	}

	// O AutoMigrate agora criará as tabelas no Postgres do Neon
	database.AutoMigrate(&Task{})
    // No futuro, adicionaremos aqui: database.AutoMigrate(&Credential{}, &Note{})

	DB = database
}