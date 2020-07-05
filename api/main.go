package main

import (

	echo "github.com/labstack/echo/v4"
	"github.com/yongwoon/echo-blog/route"
	"github.com/yongwoon/echo-blog/config/initializer"
)

func main() {
	initializer.InitDotenv()

	e := echo.New()
	route.Init(e)

	e.Logger.Fatal(e.Start(":1323"))
}
