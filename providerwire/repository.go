//go:build wireinject
// +build wireinject

package providerwire

import (
	"github.com/google/wire"
	repositoryInterface "github.com/sonymuhamad/todo-app/pkg/interfaces/repository"
	"github.com/sonymuhamad/todo-app/repository"
)

var repositorySet = wire.NewSet(repository.NewPost)

var provideRepository = wire.NewSet(
	repositorySet,
	wire.Bind(new(repositoryInterface.PostRepository), new(*repository.Post)),
)
