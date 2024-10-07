package httpserver

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sonymuhamad/todo-app/pkg"
	"net/http"
)

func InitRouter(cfg pkg.EnvConfig, h *Handler) *echo.Echo {
	r := echo.New()
	r.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},                                                                                    // Allow all origins
		AllowMethods: []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete, http.MethodOptions}, // Specify allowed methods
		AllowHeaders: []string{"Content-Type"},                                                                         // Specify allowed headers
	}))

	r.GET("/posts", h.GetPosts)
	r.POST("/posts", h.CreatePost)
	r.PUT("/posts/:id", h.UpdatePost)
	r.DELETE("/posts/:id", h.DeletePost)

	return r
}
