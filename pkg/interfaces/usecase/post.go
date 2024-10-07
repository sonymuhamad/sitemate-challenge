package usecase

import (
	"context"
)

type PostUsecase interface {
	Create(ctx context.Context) error
}
