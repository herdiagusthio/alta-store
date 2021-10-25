package api

import (
	"altaStore/api/middleware"
	"altaStore/api/v1/admins"

	"github.com/labstack/echo/v4"
)

func RegisterPath(e *echo.Echo, adminController *admins.Controller) {
	e.POST("admins/login", adminController.LoginController)
	e.POST("adminmockdata", adminController.CreateAdminController)
	admin := e.Group("admins")
	admin.Use(middleware.JWTMiddleware())
	admin.GET("", adminController.GetAdminController)
	admin.GET("/:username", adminController.GetAdminByUsername)
	admin.POST("", adminController.CreateAdminController)
	admin.PUT("/:username", adminController.ModifyAdminController)
	//health check
	e.GET("/health", func(c echo.Context) error {
		return c.NoContent(200)
	})
}
