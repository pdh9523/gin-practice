package dto

type PostUpdateDto struct {
	Title   *string `json:"title,omitempty"`
	Content *string `json:"body,omitempty"`
}
