package repository

import (
	"context"
	"github.com/sonymuhamad/todo-app/model"
)

type PostRepository interface {
	GetByID(ctx context.Context, ID int64) (model.Post, error)
}
