package route

import (
	echo "github.com/labstack/echo/v4"
)

// Init Initialiize routes
func Init(e *echo.Echo) {

	e.GET("/ping", func(c echo.Context) error {
		return c.String(200, "OK")
	})

	routeV1(e)
}
