package models

import (
	
	"gorm.io/gorm"
)


type Study struct {
    gorm.Model
    Title       string `json:"title"`
    Link        string `json:"link"`       // Para guardar um link do Notion ou PDF
    Status      string `json:"status"`     // "Estudando", "Lido", "Pausado"
    UserID      uint   `json:"user_id"`    // Relacionamento com o usuário
    Focos   string `json:"focos"`  // Guardaremos os focos aqui
}