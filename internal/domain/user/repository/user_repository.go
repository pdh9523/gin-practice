package repository

import "github.com/pdh9523/gin-practice/internal/domain/user/model"

type UserRepository interface {
	Create(user *model.User) error
	FindByID(ID uint) (*model.User, error)
	FindByEmail(email string) (*model.User, error)
}
