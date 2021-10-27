package request

import (
	"altaStore/business/address"
)

type InsertAddressRequest struct {
	UserID      uint   `validate:"required"`
	Name        string `validate:"required"`
	PhoneNumber string `validate:"required,number"`
	Street      string `validate:"required"`
	City        string `validate:"required"`
	Province    string `validate:"required"`
	District    string `validate:"required"`
	PostalCode  uint   `validate:"required"`
	AddressType string
	IsDefault   bool
}

func (req *InsertAddressRequest) ToUpsertAddressSpec() *address.InsertAddressSpec {

	var insertAddressSpec address.InsertAddressSpec

	insertAddressSpec.UserID = req.UserID
	insertAddressSpec.Name = req.Name
	insertAddressSpec.PhoneNumber = req.PhoneNumber
	insertAddressSpec.Street = req.Street
	insertAddressSpec.City = req.City
	insertAddressSpec.Province = req.Province
	insertAddressSpec.District = req.District
	insertAddressSpec.PostalCode = req.PostalCode
	insertAddressSpec.AddressType = &req.AddressType
	insertAddressSpec.IsDefault = req.IsDefault

	return &insertAddressSpec
}
