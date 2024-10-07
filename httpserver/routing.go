package httpserver

import (
	"github.com/labstack/echo/v4"
	"github.com/sonymuhamad/todo-app/pkg"
)

func InitRouter(cfg pkg.EnvConfig, h *Handler) *echo.Echo {
	r := echo.New()

	r.GET("/posts", h.GetPosts)
	r.POST("/posts", h.CreatePost)
	r.PUT("/posts/:id", h.UpdatePost)
	r.DELETE("/posts/:id", h.DeletePost)

	return r
}
