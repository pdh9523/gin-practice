package service

import (
	"github.com/pdh9523/gin-practice/internal/domain/user/dto"
	"github.com/pdh9523/gin-practice/internal/domain/user/repository"
	"golang.org/x/crypto/bcrypt"
)

type UserServiceImpl struct {
	UserRepository repository.UserRepository
	PreUserStore   repository.PreUserStore
}

func NewUserService(userRepository repository.UserRepository, preUserStore repository.PreUserStore) UserService {
	return &UserServiceImpl{UserRepository: userRepository, PreUserStore: preUserStore}
}

func (s *UserServiceImpl) RegisterUser(userRequestDto dto.UserRequestDto) (*string, error) {
	user := dto.ToUser(userRequestDto)

	hashed, err := bcrypt.GenerateFromPassword([]byte(userRequestDto.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user.Password = string(hashed)

	err = s.PreUserStore.Save(user)
	// 임시 회원이 된 유저의 이메일을 리턴해서, 바로 메일을 보낼 수 있도록 처리
	return &user.Email, err
}
