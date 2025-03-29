package service

import (
	"github.com/pdh9523/gin-practice/internal/domain/user/dto"
	"github.com/pdh9523/gin-practice/internal/domain/user/model"
	"github.com/pdh9523/gin-practice/internal/domain/user/repository"
	"golang.org/x/crypto/bcrypt"
)

type UserServiceImpl struct {
	Repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &UserServiceImpl{Repo: repo}
}

func (s *UserServiceImpl) RegisterUser(userRequestDto dto.UserRequestDto) (*model.User, error) {
	user := dto.ToUser(userRequestDto)

	hashed, err := bcrypt.GenerateFromPassword([]byte(userRequestDto.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user.Password = string(hashed)
	if err := s.Repo.Create(user); err != nil {
		return nil, err
	}

	return user, nil
}
