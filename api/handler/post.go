package handler

import (
	"net/http"
	"strconv"

	echo "github.com/labstack/echo/v4"
	"github.com/yongwoon/echo-sqlboiler-mysql-blog/service"
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

	postCreateReq struct {
		Title string `json:"title"`
		Body  string `json:"body"`
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

	res, err := service.PostByID(id)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, res)
}

// Create : POST api/vx/posts
func (p *Post) Create(c echo.Context) error {
	pcr := &postCreateReq{}
	if err := c.Bind(pcr); err != nil {
		return err
	}

	post, err := service.PostCreate(pcr.Title, pcr.Body)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, post)
}
