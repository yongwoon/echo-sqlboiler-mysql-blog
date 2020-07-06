package route

import (
	echo "github.com/labstack/echo/v4"
	"github.com/yongwoon/echo-blog/config"
	"github.com/yongwoon/echo-blog/handler"
)

func routeV1(e *echo.Echo) {
	v1 := e.Group(config.APIVersion1)
	posts := v1.Group("/posts")

	postHandler := handler.NewPost()
	posts.GET("", postHandler.Index)
	posts.GET("/:id", postHandler.Show)
	posts.POST("", postHandler.Create)
}
