package repository

import "github.com/pdh9523/gin-practice/internal/domain/user/model"

type PreUserStore interface {
	Save(user *model.User) error
	FindByEmail(email string) (*model.User, error)
}
