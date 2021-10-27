package response

import (
	"altaStore/business/address"
)

type getAllAddressResponse struct {
	Addresses []GetAddressResponse `json:"addresses"`
}

//NewGetAllUserResponse construct GetAllUserResponse
func NewGetAllAddressResponse(addresses []address.Address) getAllAddressResponse {
	getAllAddressResponse := getAllAddressResponse{}
	for _, value := range addresses {

		var getAddressResponse GetAddressResponse

		getAddressResponse.ID = value.ID
		getAddressResponse.UserID = value.UserID
		getAddressResponse.Name = value.Name
		getAddressResponse.PhoneNumber = value.PhoneNumber
		getAddressResponse.Street = value.Street
		getAddressResponse.City = value.City
		getAddressResponse.Province = value.Province
		getAddressResponse.District = value.District
		getAddressResponse.PostalCode = value.PostalCode
		getAddressResponse.AddressType = value.AddressType
		getAddressResponse.IsDefault = value.IsDefault
		getAddressResponse.CreatedAt = value.CreatedAt
		getAddressResponse.UpdatedAt = value.UpdatedAt
		getAddressResponse.DeletedAt = value.DeletedAt

		getAllAddressResponse.Addresses = append(getAllAddressResponse.Addresses, getAddressResponse)
	}

	if getAllAddressResponse.Addresses == nil {
		getAllAddressResponse.Addresses = []GetAddressResponse{}
	}

	return getAllAddressResponse
}
