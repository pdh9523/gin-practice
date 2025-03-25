package service

import (
	"github.com/pdh9523/gin-practice/internal/domain/post/dto"
	"github.com/pdh9523/gin-practice/internal/domain/post/model"
)

type PostService interface {
	GetPosts() ([]*model.Post, error)
	GetPostByID(id uint) (*model.Post, error)
	CreatePost(postRequestDto dto.PostRequestDto) (*model.Post, error)
	UpdatePost(id uint, postUpdateDto dto.PostUpdateDto) (*model.Post, error)
	DeletePost(id uint) error
}
