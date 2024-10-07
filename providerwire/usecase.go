//go:build wireinject
// +build wireinject

package providerwire

import (
	"github.com/google/wire"
	usecaseInterface "github.com/sonymuhamad/todo-app/pkg/interfaces/usecase"
	"github.com/sonymuhamad/todo-app/usecase"
)

var usecaseSet = wire.NewSet(usecase.NewPost)

var provideUsecase = wire.NewSet(
	usecaseSet,
	wire.Bind(new(usecaseInterface.PostUsecase), new(*usecase.Post)),
)
