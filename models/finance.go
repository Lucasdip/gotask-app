package models 

import (
    "gorm.io/gorm"
)



type Transaction struct {
    gorm.Model
    Description string  `json:"description"`
    Amount      float64 `json:"amount"`
    Type        string  `json:"type"` // "income" (ganho) ou "expense" (gasto)
    UserID      uint    `json:"user_id"`
}