package auth

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// Routes defines the URL routing for Authentication middleware.
func Routes(e *echo.Echo) *echo.Group {
	r := e.Group("/admin")
	r.Use(middleware.JWT([]byte("secret")))

	r.GET("", Restricted)

	e.POST("/login", Login)
	e.GET("/", Accessible)

	return r
}
