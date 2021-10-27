package api

import (
	"altaStore/api/middleware"
	"altaStore/api/v1/address"
	"altaStore/api/v1/auth"
	"altaStore/api/v1/cart"
	"altaStore/api/v1/user"

	echo "github.com/labstack/echo/v4"
)

//RegisterPath Register all API with routing path
func RegisterPath(
	e *echo.Echo,
	authController *auth.Controller,
	userController *user.Controller,
	addressController *address.Controller,
	cartController *cart.Controller) {
	if authController == nil || userController == nil || addressController == nil || cartController == nil {
		panic("Controller parameter cannot be nil")
	}

	// 	//authentication with Versioning endpoint
	authV1 := e.Group("v1/auth")
	authV1.POST("/login", authController.Login)

	//user with Versioning endpoint
	e.POST("v1/users", userController.InsertUser)
	userV1 := e.Group("v1/users")
	userV1.Use(middleware.JWTMiddleware())
	userV1.GET("/:id", userController.FindUserByID)
	userV1.GET("", userController.FindAllUser)
	userV1.POST("/address", addressController.InsertAddress)
	userV1.GET("/address", addressController.GetAllAddress)
	userV1.GET("/address/default", addressController.GetDefaultAddress)

	cartV1 := e.Group("v1/cart")
	cartV1.Use(middleware.JWTMiddleware())
	cartV1.POST("", cartController.AddToCart)

	// 	userV1.PUT("/:id", userController.UpdateUser)

	// 	//pet with Versioning endpoint
	// 	petV1 := e.Group("v1/pets")
	// 	petV1.Use(middleware.JWTMiddleware())
	// 	petV1.GET("/:id", petController.FindPetByID)
	// 	petV1.GET("", petController.FindAllPet)
	// 	petV1.POST("", petController.InsertPet)
	// 	petV1.PUT("/:id", petController.UpdatePet)

	// 	//health check
	// 	e.GET("/health", func(c echo.Context) error {
	// 		return c.NoContent(200)
	// 	})
}
