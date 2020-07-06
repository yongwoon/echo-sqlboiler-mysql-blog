package handler

import (
	"net/http"
	"strconv"

	echo "github.com/labstack/echo/v4"
	"github.com/yongwoon/echo-blog/service"
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

	res, err := service.PostList()
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, res)
}

// Show : api/vX/posts/:id
func (p *Post) Show(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)

	res, err := service.PostById(id)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, res)
}
