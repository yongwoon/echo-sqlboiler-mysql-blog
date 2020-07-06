package handler

import (
	"net/http"

	echo "github.com/labstack/echo/v4"
)

type (
	// IPost interface
	IPost interface {
		Index(echo.Context) error
		Show(echo.Context) error
	}
	// Post struct
	Post struct {
	}
)

// NewPost returns NewPost instance.
func NewPost() *Post {
	return &Post{}
}

// Index : api/vX/posts
func (p *Post) Index(c echo.Context) error {
	return c.JSON(http.StatusOK, "posts index")
}

// Show : api/vX/posts/:id
func (p *Post) Show(c echo.Context) error {
	return c.JSON(http.StatusOK, "posts show")
}
