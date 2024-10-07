package repository

import (
	"context"
	"github.com/sonymuhamad/todo-app/model"
	"gorm.io/gorm"
)

type Post struct {
	BaseRepository
}

func NewPost(db *gorm.DB) *Post {
	return &Post{
		BaseRepository{
			db: db,
		},
	}
}

func (p *Post) GetByID(ctx context.Context, ID int64) (model.Post, error) {
	// example repository for get post by ID
	return model.Post{}, nil
}
