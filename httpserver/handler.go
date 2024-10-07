package httpserver

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/sonymuhamad/todo-app/model"
	usecaseInterface "github.com/sonymuhamad/todo-app/pkg/interfaces/usecase"
	"net/http"
	"strconv"
)

type Handler struct {
	postUsecase usecaseInterface.PostUsecase
}

func NewHandlerWithWire(postUsecase usecaseInterface.PostUsecase) *Handler {
	return &Handler{
		postUsecase: postUsecase,
	}
}

// Mock data
var mockPosts = []model.Post{
	{ID: 1, Title: "First Post", Description: "This is the body of the first post."},
	{ID: 2, Title: "Second Post", Description: "This is the body of the second post."},
	{ID: 3, Title: "Third Post", Description: "This is the body of the third post."},
}

func (h *Handler) GetPosts(c echo.Context) error {
	return c.JSON(http.StatusOK, mockPosts)
}

func (h *Handler) CreatePost(c echo.Context) error {
	post := new(model.Post)
	if err := c.Bind(post); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	post.ID = int64(len(mockPosts) + 1)

	mockPosts = append(mockPosts, *post)

	// print post to console
	fmt.Println(post)

	return c.JSON(http.StatusOK, post)
}

func (h *Handler) UpdatePost(c echo.Context) error {
	// Get the post ID from the URL parameter
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid post ID"})
	}

	// Find the post by ID
	var post *model.Post
	for i := range mockPosts {
		if mockPosts[i].ID == int64(id) {
			post = &mockPosts[i]
			break
		}
	}

	if post == nil {
		return c.JSON(http.StatusNotFound, map[string]string{"message": "Post not found"})
	}

	// Bind the updated data to the post
	if err := c.Bind(post); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid request body"})
	}

	// print post to console
	fmt.Println(post)

	return c.JSON(http.StatusOK, post)
}

func (h *Handler) DeletePost(c echo.Context) error {
	// Get the post ID from the URL parameter
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid post ID"})
	}

	// Find the index of the post by ID
	index := -1
	for i, post := range mockPosts {
		if post.ID == int64(id) {
			index = i
			break
		}
	}

	if index == -1 {
		return c.JSON(http.StatusNotFound, map[string]string{"message": "Post not found"})
	}

	// Delete the post from the slice
	mockPosts = append(mockPosts[:index], mockPosts[index+1:]...)

	return c.JSON(http.StatusOK, map[string]string{"message": "Post deleted successfully"})
}
