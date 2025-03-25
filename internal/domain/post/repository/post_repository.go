package repository

import "github.com/pdh9523/gin-practice/internal/domain/post/model"

type PostRepository interface {
	FindAll() ([]*model.Post, error)
	FindByID(id uint) (*model.Post, error)
	Create(post *model.Post) error
	Update(post *model.Post) error
	DeleteByID(id uint) error
}
