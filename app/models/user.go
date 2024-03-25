package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID       uuid.UUID `gorm:"column:id; gorm:"type:uuid;default:uuid_generate_v4()" json:"id"`
	Name     string    `gorm:"column:name;comment:''" json:"name"`
	Username string    `gorm:"column:username;comment:''" json:"username"`
	Password string    `gorm:"column:password;comment:''" json:"password"`
}

type Create struct {
	Name     string `gorm:"column:name;comment:''" json:"name"`
	Username string `gorm:"column:username;comment:''" json:"username"`
	Password string `gorm:"column:password;comment:''" json:"password"`
}

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
