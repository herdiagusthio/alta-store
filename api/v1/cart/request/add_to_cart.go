package request

import "altaStore/business/cart"

type AddToCartRequest struct {
	UserID    uint `validate:"required"`
	ProductID uint `validate:"required"`
	Price     uint `validate:"required"`
	Quantity  uint `validate:"required"`
}

func (req *AddToCartRequest) AddToCartSpec() *cart.AddToCartSpec {

	var addToCartSpec cart.AddToCartSpec

	addToCartSpec.UserID = req.UserID
	addToCartSpec.ProductID = req.ProductID
	addToCartSpec.Price = req.Price
	addToCartSpec.Quantity = req.Quantity

	return &addToCartSpec
}
