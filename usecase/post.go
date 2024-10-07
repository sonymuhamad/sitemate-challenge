package usecase

import (
	"context"
	"github.com/sonymuhamad/todo-app/pkg/interfaces/repository"
)

type Post struct {
	postRepository repository.PostRepository
}

func NewPost(postRepository repository.PostRepository) *Post {
	return &Post{
		postRepository: postRepository,
	}
}

func (p *Post) Create(ctx context.Context) error {
	// example usecase for create post
	return nil
}
