package admins

import (
	"altaStore/business"
	"altaStore/util/validator"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type AdminSpec struct {
	Name         string `form:"name" validate:"required,max=20"`
	Password     string `form:"password" validate:"required"`
	Email        string `form:"email" validate:"required,email"`
	Phone_number string `form:"phone_number" validate:"required,max=20"`
	Username     string `form:"username" validate:"required,max=10"`
	CreatedBy    string `form:"created_by" validate:"required"`
}

type AdminUpdatable struct {
	Name         string `json:"name" form:"name" validate:"required,max=20"`
	Phone_number string `json:"phone_number" form:"phone_number" validate:"required,max=20"`
	Status       string `json:"status" form:"status" validate:"max=10"`
	Modified_by  string `json:"modified_by" form:"modified_by" validate:"required"`
}

type service struct {
	AdminsRepository Repository
}

func InitAdminService(repository Repository) *service {
	return &service{
		AdminsRepository: repository,
	}
}

func (admin_service *service) FindAllAdmin(offset, limit int) (*[]Admins, error) {
	list_admins, err := admin_service.AdminsRepository.GetAdmin(limit, offset)
	if err != nil {
		return nil, err
	}
	return list_admins, err
}

func (admin_service *service) FindAdminByUsername(username string) (*Admins, error) {
	admins, err := admin_service.AdminsRepository.GetAdminByUsername(username)
	if err != nil {
		return nil, err
	}
	return admins, err
}

func (admin_service *service) FindAdminById(id_admins int) (*Admins, error) {
	admins, err := admin_service.AdminsRepository.GetAdminById(id_admins)
	if err != nil {
		return nil, err
	}
	return admins, err
}

func (admin_service *service) LoginAdmin(username, password string) (*Admins, error) {
	// hashing data passwords
	admin, err := admin_service.AdminsRepository.LoginAdmin(username, password)
	if err != nil {
		return nil, err
	}
	return admin, nil
}

func (admin_service *service) InsertAdmin(admin_spec AdminSpec, createdBy string) error {
	err := validator.GetValidator().Struct(admin_spec)
	if err != nil {
		return business.ErrInvalidSpec
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(admin_spec.Password), 0)

	if err != nil {
		return err
	}

	admin := NewAdmin(admin_spec.Name,
		admin_spec.Username,
		admin_spec.Email,
		string(hashedPassword),
		admin_spec.Phone_number,
		admin_spec.CreatedBy,
	)

	err = admin_service.AdminsRepository.CreateAdmin(&admin)
	if err != nil {
		return err
	}
	return nil
}

func (admin_service *service) ModifyAdmin(username string, admin_updatable AdminUpdatable) error {
	err := validator.GetValidator().Struct(admin_updatable)
	if err != nil {
		return err
	}
	fmt.Println(username)
	admin_data, err := admin_service.AdminsRepository.GetAdminByUsername(username)
	if err != nil {
		return err
	}

	fmt.Println(admin_data)
	new_admin_data := admin_data.ModifyAdmin(admin_updatable)
	err = admin_service.AdminsRepository.UpdateAdmin(&new_admin_data)
	if err != nil {
		return err
	}
	return nil
}
