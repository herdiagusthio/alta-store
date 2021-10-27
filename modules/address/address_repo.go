package address

import (
	"altaStore/business/address"
	"time"

	"gorm.io/gorm"
)

type GormRepository struct {
	DB *gorm.DB
}

type Address struct {
	ID          uint    `gorm:"id;primaryKey;autoIncrement"`
	UserID      uint    `json:"user_id" validate:"required" gorm:"type:integer; not null"`
	Name        string  `json:"name" validate:"required" gorm:"type:varchar(100); not null"`
	PhoneNumber string  `json:"phone_number" validate:"required,number" gorm:"type:varchar(20); not null"`
	Street      string  `json:"street" validate:"required" gorm:"type:text; not null"`
	City        string  `json:"city" validate:"required" gorm:"type:varchar(50); not null"`
	Province    string  `json:"province" validate:"required" gorm:"type:varchar(50); not null"`
	District    string  `json:"district" validate:"required" gorm:"type:varchar(50); not null"`
	PostalCode  uint    `json:"postal_code" validate:"required" gorm:"not null"`
	AddressType *string `json:"address_type" gorm:"type:varchar(50)"`
	IsDefault   bool    `json:"is_default" gorm:"default:false"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time
}

func newAddressTable(address address.Address) *Address {

	return &Address{
		address.ID,
		address.UserID,
		address.Name,
		address.PhoneNumber,
		address.Street,
		address.City,
		address.Province,
		address.District,
		address.PostalCode,
		address.AddressType,
		address.IsDefault,
		address.CreatedAt,
		address.UpdatedAt,
		address.DeletedAt,
	}

}

func (col *Address) ToAddress() address.Address {
	var address address.Address

	address.ID = col.ID
	address.UserID = col.UserID
	address.Name = col.Name
	address.PhoneNumber = col.PhoneNumber
	address.Street = col.Street
	address.City = col.City
	address.Province = col.Province
	address.District = col.District
	address.PostalCode = col.PostalCode
	address.AddressType = col.AddressType
	address.IsDefault = col.IsDefault
	address.CreatedAt = col.CreatedAt
	address.UpdatedAt = col.UpdatedAt
	address.DeletedAt = col.DeletedAt

	return address
}

//NewGormDBRepository Generate Gorm DB user repository
func NewGormDBRepository(db *gorm.DB) *GormRepository {
	return &GormRepository{
		db,
	}
}

//InsertAddress Insert new Address into storage
func (repo *GormRepository) InsertAddress(address address.Address) error {

	addressData := newAddressTable(address)
	addressData.ID = 0

	err := repo.DB.Create(addressData).Error
	if err != nil {
		return err
	}

	return nil
}

//GetDefaultAddress Get default address from storage
func (repo *GormRepository) GetDefaultAddress(userID uint) (*address.Address, error) {
	var userAddress Address

	err := repo.DB.Where("user_id = ?", userID).Where("is_default = ?", true).First(&userAddress).Error
	if err != nil {
		return nil, err
	}

	address := userAddress.ToAddress()

	return &address, nil
}

//UpdateDefaultAddress Update is_default coloum to false
func (repo *GormRepository) UpdateDefaultAddress(address address.Address) error {
	addressData := newAddressTable(address)
	err := repo.DB.Model(&addressData).Update("is_default", false).Error
	if err != nil {
		return err
	}

	return nil
}

func (repo *GormRepository) GetAllAddress(userID uint) ([]address.Address, error) {
	var addresses []Address

	err := repo.DB.Where("user_id = ?", userID).Find(&addresses).Error
	if err != nil {
		return nil, err
	}

	var result []address.Address
	for _, value := range addresses {
		result = append(result, value.ToAddress())
	}

	return result, nil
}
