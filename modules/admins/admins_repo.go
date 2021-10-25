package admins

import (
	"altaStore/business/admins"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type GormRepository struct {
	DB *gorm.DB
}

type AdminsTable struct {
	gorm.Model
	ID           int       `gorm:"id;primaryKey:autoIncrement"`
	Name         string    `gorm:"name;type:varchar(20);not null"`
	Status       string    `gorm:"status;type:varchar(10);not null"`
	Password     string    `gorm:"password;not null"`
	Email        string    `gorm:"email;type:varchar(30);unique;not null"`
	Phone_number string    `gorm:"phone_number;type:varchar(20);not null"`
	Username     string    `gorm:"username;type:varchar(10);uniqueIndex:username,sort:asc;not null"`
	Created_by   string    `gorm:"created_by"`
	Modified_by  string    `gorm:"modified_by"`
	Created_at   time.Time `gorm:"created_at;type:datetime;default:null"`
	Updated_at   time.Time `gorm:"updated_at;type:datetime;default:null"`
	Deleted_at   time.Time `gorm:"deleted_at;type:datetime;default:null"`
}

func ConvertToAdminTable(admin *admins.Admins) *AdminsTable {
	return &AdminsTable{
		Name:         admin.Name,
		Status:       admin.Status,
		Password:     admin.Password,
		Email:        admin.Email,
		Phone_number: admin.Phone_number,
		Username:     admin.Username,
		Created_by:   admin.CreatedBy,
		Created_at:   admin.Created_at,
		Updated_at:   admin.Updated_at,
		Deleted_at:   admin.Deleted_at,
		Modified_by:  admin.Modified_by,
	}
}

func (admins_table *AdminsTable) ConvertToAdmin() *admins.Admins {
	return &admins.Admins{
		ID:           admins_table.ID,
		Name:         admins_table.Name,
		Status:       admins_table.Status,
		Email:        admins_table.Email,
		Phone_number: admins_table.Phone_number,
		Username:     admins_table.Username,
		Created_at:   admins_table.Created_at,
		Updated_at:   admins_table.UpdatedAt,
		CreatedBy:    admins_table.Created_by,
		Deleted_at:   admins_table.UpdatedAt,
		Modified_by:  admins_table.Modified_by,
	}
}

func InitAdminRepository(db *gorm.DB) *GormRepository {
	return &GormRepository{
		DB: db,
	}
}

func (repository *GormRepository) GetAdmin(limit, offset int) (*[]admins.Admins, error) {
	// should be using skip n offset
	var list_admin_tables []AdminsTable
	err := repository.DB.Offset(offset - 1).Limit(limit).Find(&list_admin_tables).Error
	if err != nil {
		return nil, err
	}
	var list_admin []admins.Admins
	for _, admin_tables := range list_admin_tables {
		list_admin = append(list_admin, *admin_tables.ConvertToAdmin())
	}
	return &list_admin, nil
}

func (repository *GormRepository) GetAdminByUsername(username string) (*admins.Admins, error) {
	admin := AdminsTable{}
	err := repository.DB.Where("username = ?", username).First(&admin).Error
	if err != nil {
		return nil, err
	}
	return admin.ConvertToAdmin(), nil
}

func (repository *GormRepository) GetAdminById(id int) (*admins.Admins, error) {
	admin := AdminsTable{}
	err := repository.DB.Where("id = ?", id).First(&admin).Error
	if err != nil {
		return nil, err
	}
	return admin.ConvertToAdmin(), nil
}

func (repository *GormRepository) LoginAdmin(username, password string) (*admins.Admins, error) {
	admin := AdminsTable{}
	err := repository.DB.Where("username = ? ", username).First(&admin).Error
	if err != nil {
		return nil, err
	}

	if errCompare := bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(password)); errCompare != nil {
		return nil, errCompare
	}
	return admin.ConvertToAdmin(), nil
}

func (repository *GormRepository) CreateAdmin(admin *admins.Admins) error {
	admin_table := ConvertToAdminTable(admin)
	err := repository.DB.Save(admin_table).Error
	if err != nil {
		return err
	}
	return nil
}

func (repository *GormRepository) UpdateAdmin(admin *admins.Admins) error {
	admin_table := ConvertToAdminTable(admin)
	err := repository.DB.Where("username = ?", admin.Username).Model(admin_table).Updates(AdminsTable{
		Name:         admin.Name,
		Status:       admin.Status,
		Phone_number: admin.Phone_number,
		Modified_by:  admin.Modified_by,
		Updated_at:   admin.Updated_at,
	}).Error
	if err != nil {
		return err
	}
	return nil
}
