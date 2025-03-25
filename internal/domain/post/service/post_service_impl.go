package service

import (
	"github.com/pdh9523/gin-practice/internal/domain/post/dto"
	"github.com/pdh9523/gin-practice/internal/domain/post/model"
	"github.com/pdh9523/gin-practice/internal/domain/post/repository"
)

type PostServiceImpl struct {
	Repository repository.PostRepository
}

func NewPostService(repository repository.PostRepository) PostService {
	return &PostServiceImpl{Repository: repository}
}

func (s *PostServiceImpl) GetPosts() ([]*model.Post, error) {
	posts, err := s.Repository.FindAll()
	if err != nil {
		return nil, err
	}
	return posts, nil
}

func (s *PostServiceImpl) GetPostByID(id uint) (*model.Post, error) {
	post, err := s.Repository.FindByID(id)
	if err != nil {
		return nil, err
	}
	return post, nil
}

func (s *PostServiceImpl) CreatePost(postRequestDto dto.PostRequestDto) (*model.Post, error) {
	post := dto.ToPost(postRequestDto)

	if err := s.Repository.Create(post); err != nil {
		return nil, err
	}
	return post, nil
}

func (s *PostServiceImpl) UpdatePost(id uint, postUpdateDto dto.PostUpdateDto) (*model.Post, error) {
	post, err := s.GetPostByID(id)
	if err != nil {
		return nil, err
	}

	if postUpdateDto.Title != nil {
		post.Title = *postUpdateDto.Title
	}
	if postUpdateDto.Content != nil {
		post.Content = *postUpdateDto.Content
	}

	err = s.Repository.Update(post)
	return post, err
}

func (s *PostServiceImpl) DeletePost(id uint) error {
	return s.Repository.DeleteByID(id)
}
