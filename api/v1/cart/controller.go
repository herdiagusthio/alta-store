package cart

import (
	"altaStore/api/common"
	"altaStore/api/v1/cart/request"
	"altaStore/business/cart"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

//Controller Get item API controller
type Controller struct {
	service cart.Service
}

//NewController Construct item API controller
func NewController(service cart.Service) *Controller {
	return &Controller{
		service,
	}
}

// InsertAddress Create new address handler
func (controller *Controller) AddToCart(c echo.Context) error {

	user := c.Get("user").(*jwt.Token)
	if !user.Valid {
		return c.JSON(common.NewForbiddenResponse())
	}

	claims := user.Claims.(jwt.MapClaims)

	//use float64 because its default data that provide by JWT, we cant use int/int32/int64/etc.
	//MUST CONVERT TO FLOAT64, OTHERWISE ERROR (not _ok_)!
	userID, ok := claims["id"].(float64)
	if !ok {
		return c.JSON(common.NewForbiddenResponse())
	}

	addToCartRequest := new(request.AddToCartRequest)
	addToCartRequest.UserID = uint(userID)

	if err := c.Bind(addToCartRequest); err != nil {
		return c.JSON(common.NewBadRequestResponse())
	}

	err := controller.service.AddToCart(*addToCartRequest.AddToCartSpec())
	if err != nil {
		return c.JSON(common.NewErrorBusinessResponse(err))
	}

	return c.JSON(common.NewSuccessResponseWithoutData())
}
