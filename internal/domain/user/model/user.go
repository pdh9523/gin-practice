package model

import (
	"github.com/pdh9523/gin-practice/internal/domain/post/model"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	Nickname string       `json:"nickname" gorm:"not null;unique"`
	Email    string       `json:"email" gorm:"not null;unique"`
	Password string       `json:"password" gorm:"not null"`
	Posts    []model.Post `gorm:"foreignkey:UserID"`
}
