package api

import "github.com/labstack/echo/v4"

// Login API for user
func UserLogin(e *echo.Echo) {
	e.GET("/auth/user/login", func(c echo.Context) error {
		return c.String(200, "User Login API")
	})
}
