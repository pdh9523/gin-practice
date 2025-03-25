package repository

import (
	"github.com/pdh9523/gin-practice/internal/domain/post/model"
	"gorm.io/gorm"
)

type GormPostRepository struct {
	DB *gorm.DB
}

func NewGormPostRepository(DB *gorm.DB) PostRepository {
	return &GormPostRepository{DB: DB}
}

func (r *GormPostRepository) FindAll() ([]*model.Post, error) {
	posts := make([]*model.Post, 0)
	err := r.DB.Find(&posts).Error
	return posts, err
}

func (r *GormPostRepository) FindByID(id uint) (*model.Post, error) {
	post := &model.Post{}
	err := r.DB.Where("id = ?", id).First(post).Error
	return post, err
}

func (r *GormPostRepository) Create(post *model.Post) error {
	return r.DB.Create(post).Error
}

func (r *GormPostRepository) Update(post *model.Post) error {
	return r.DB.Save(post).Error
}

func (r *GormPostRepository) DeleteByID(id uint) error {
	return r.DB.Delete(&model.Post{}, id).Error
}
