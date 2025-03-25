package dto

import "github.com/pdh9523/gin-practice/internal/domain/post/model"

type PostRequestDto struct {
	Title   string `json:"title" validate:"required"`
	Content string `json:"body" validate:"required"`
}

func ToPost(req PostRequestDto) *model.Post {
	return &model.Post{
		Title:   req.Title,
		Content: req.Content,
	}
}
