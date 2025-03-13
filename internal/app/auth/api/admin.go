package api

import "github.com/labstack/echo/v4"

// Login API for admin
func AdminLogin(e *echo.Echo) {
	e.GET("/auth/admin/login", func(c echo.Context) error {
		return c.String(200, "Admin Login API")
	})
}
