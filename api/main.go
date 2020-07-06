package main

import (
	"github.com/labstack/echo/v4"
	"github.com/yongwoon/echo-sqlboiler-mysql-blog/initializer"
	"github.com/yongwoon/echo-sqlboiler-mysql-blog/route"
)

func main() {
	initializer.InitDotenv()

	e := echo.New()
	route.Init(e)

	e.Logger.Fatal(e.Start(":1323"))
}
