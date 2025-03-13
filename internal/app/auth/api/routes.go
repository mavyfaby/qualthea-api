package api

import "github.com/labstack/echo/v4"

// Register routes for the Authentication API
func RegisterRoutes(server *echo.Echo) {
	UserLogin(server)
	AdminLogin(server)
}
