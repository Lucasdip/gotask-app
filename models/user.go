package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"unique;not null" json:"username"`
	Password string `gorm:"not null" json:"password"` // Em produção, use Hash (bcrypt)
	XP       int    `json:"xp" gorm:"default:0"`
    Level    int    `json:"level" gorm:"default:1"`
}