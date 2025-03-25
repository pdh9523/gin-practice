package repository

import (
	"github.com/pdh9523/gin-practice/internal/domain/user/model"
	"gorm.io/gorm"
)

type GormUserRepository struct {
	DB *gorm.DB
}

func NewGormUserRepository(db *gorm.DB) UserRepository {
	return &GormUserRepository{DB: db}
}

func (r *GormUserRepository) Create(user *model.User) error {
	return r.DB.Create(user).Error
}

func (r *GormUserRepository) FindByID(ID uint) (*model.User, error) {
	var user model.User
	err := r.DB.Where("id=?", ID).First(&user).Error
	return &user, err
}

func (r *GormUserRepository) FindByEmail(email string) (*model.User, error) {
	var user model.User
	err := r.DB.Where("email=?", email).First(&user).Error
	return &user, err
}
