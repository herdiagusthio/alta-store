package address

import (
	"altaStore/api/common"
	"altaStore/api/v1/address/request"
	"altaStore/api/v1/address/response"
	"altaStore/business/address"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

//Controller Get item API controller
type Controller struct {
	service address.Service
}

//NewController Construct item API controller
func NewController(service address.Service) *Controller {
	return &Controller{
		service,
	}
}

// InsertAddress Create new address handler
func (controller *Controller) InsertAddress(c echo.Context) error {

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

	insertAddressRequest := new(request.InsertAddressRequest)
	insertAddressRequest.UserID = uint(userID)

	if err := c.Bind(insertAddressRequest); err != nil {
		return c.JSON(common.NewBadRequestResponse())
	}

	err := controller.service.InsertAddress(*insertAddressRequest.ToUpsertAddressSpec())
	if err != nil {
		return c.JSON(common.NewErrorBusinessResponse(err))
	}

	return c.JSON(common.NewSuccessResponseWithoutData())
}

//GetDefaultAddress Get default address by UserID echo handler
func (controller *Controller) GetDefaultAddress(c echo.Context) error {

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

	address, err := controller.service.GetDefaultAddress(uint(userID))
	if err != nil {
		return c.JSON(common.NewErrorBusinessResponse(err))
	}

	response := response.NewGetAddressResponse(*address)

	return c.JSON(common.NewSuccessResponse(response))
}

// GetAllAddress Get All Address
func (controller *Controller) GetAllAddress(c echo.Context) error {

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

	addresses, err := controller.service.GetAllAddress(uint(userID))
	if err != nil {
		return c.JSON(common.NewErrorBusinessResponse(err))
	}

	response := response.NewGetAllAddressResponse(addresses)

	return c.JSON(common.NewSuccessResponse(response))
}
