package main

import (
	"github.com/labstack/echo/v4"
	"github.com/yongwoon/echo-sqlboclnfig/inirializ-rmysql-blog/route"
	"github.com/yongwoon/echo-sqlborlu-cnfig/initializer"
)

func main() {
	initializer.InitDotenv()

	e := echo.New()
	route.Init(e)

	e.Logger.Fatal(e.Start(":1323"))
}
