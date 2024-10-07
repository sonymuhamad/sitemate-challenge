package httpserver

import (
	"github.com/labstack/echo/v4"
	usecaseInterface "github.com/sonymuhamad/todo-app/pkg/interfaces/usecase"
	"net/http"
)

type Handler struct {
	postUsecase usecaseInterface.PostUsecase
}

func NewHandlerWithWire(postUsecase usecaseInterface.PostUsecase) *Handler {
	return &Handler{
		postUsecase: postUsecase,
	}
}

func (h *Handler) Index(ctx echo.Context) error {
	data := "Hello from /index"
	return ctx.String(http.StatusOK, data)
}
