package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"unique;not null" json:"username"`
	Password string `gorm:"not null" json:"password"` // Em produção, use Hash (bcrypt)
	XP       int    `gorm:"default:0" json:"xp"`
    Level    int    `gorm:"default:1" json:"level"`
    Rank     string `gorm:"default:'Prisioneiro'" json:"rank"`
}