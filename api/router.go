package api

import (
	"altaStore/api/v1/user"

	"github.com/labstack/echo"
)

//RegisterPath Register all API with routing path
func RegisterPath(e *echo.Echo, userController *user.Controller) {
	if userController == nil {
		panic("Controller parameter cannot be nil")
	}

	// 	//authentication with Versioning endpoint
	// 	authV1 := e.Group("v1/auth")
	// 	authV1.POST("/login", authController.Login)

	//user with Versioning endpoint
	userV1 := e.Group("v1/users")
	userV1.GET("/:id", userController.FindUserByID)
	userV1.GET("", userController.FindAllUser)
	// 	userV1.POST("", userController.InsertUser)
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
