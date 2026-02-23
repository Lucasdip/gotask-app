package models

import (
	
	"gorm.io/gorm"
)

type Task struct {

	gorm.Model        // Adiciona campos ID, CreatedAt, UpdatedAt, DeletedAt automaticamente
	Title     string `json:"title"`
	Done      bool   `json:"done" gorm:"default:false"`
	UserID uint   `json:"user_id"`

}

