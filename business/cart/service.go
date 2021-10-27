package cart

import (
	"altaStore/business"
	"altaStore/util/validator"
	"time"
)

//AddToCartSpec create cart and cart detail spec
type AddToCartSpec struct {
	UserID    uint `validate:"required"`
	ProductID uint `validate:"required"`
	Price     uint `validate:"required"`
	Quantity  uint `validate:"required"`
}

//=============== The implementation of those interface put below =======================
type service struct {
	repository Repository
}

//NewService Construct user service object
func NewService(repository Repository) Service {
	return &service{
		repository,
	}
}

func (s *service) AddToCart(addToCartSpec AddToCartSpec) error {
	err := validator.GetValidator().Struct(addToCartSpec)
	if err != nil {
		return business.ErrInvalidSpec
	}

	getActiveCart, err := s.repository.GetActiveCart(addToCartSpec.UserID)
	if err != nil {
		cart := NewCart(
			addToCartSpec.UserID,
			"active",
			time.Now(),
		)

		err = s.repository.CreateCart(cart)
		if err != nil {
			return err
		}

		getActiveCart, _ = s.repository.GetActiveCart(addToCartSpec.UserID)
	}

	cartDetail := NewCartDetail(
		getActiveCart.ID,
		addToCartSpec.ProductID,
		addToCartSpec.Price,
		addToCartSpec.Quantity,
		time.Now(),
	)

	err = s.repository.InsertCartDetail(cartDetail)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) GetCartDetailByCartID(cartID uint) ([]CartDetail, error) {
	return []CartDetail{}, nil
}
