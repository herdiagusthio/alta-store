package user

import (
	"altaStore/business/user"
	"time"

	"gorm.io/gorm"
)

//GormRepository The implementation of user.Repository object
type GormRepository struct {
	DB *gorm.DB
}

type User struct {
	ID          uint   `gorm:"id;primaryKey;autoIncrement"`
	Name        string `json:"name"  validate:"required" gorm:"type:varchar(100); not null"`
	Email       string `json:"email" validate:"required,email" gorm:"type:varchar(50); unique; not null"`
	PhoneNumber string `json:"phone_number" validate:"required,number" gorm:"type:varchar(20); not null"`
	Password    string `json:"password"  validate:"required" gorm:"not null"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time
}

func newUserTable(user user.User) *User {

	return &User{
		user.ID,
		user.Name,
		user.Email,
		user.PhoneNumber,
		user.Password,
		user.CreatedAt,
		user.UpdatedAt,
		user.DeletedAt,
	}

}

func (col *User) ToUser() user.User {
	var user user.User

	user.ID = col.ID
	user.Name = col.Name
	user.Email = col.Email
	user.PhoneNumber = col.PhoneNumber
	user.Password = col.Password
	user.CreatedAt = col.CreatedAt
	user.UpdatedAt = col.UpdatedAt
	user.DeletedAt = col.DeletedAt

	return user
}

//NewGormDBRepository Generate Gorm DB user repository
func NewGormDBRepository(db *gorm.DB) *GormRepository {
	return &GormRepository{
		db,
	}
}

//FindUserByID If data not found will return nil without error
func (repo *GormRepository) FindUserByID(id int) (*user.User, error) {

	var userData User
	err := repo.DB.First(&userData, id).Error
	if err != nil {
		return nil, err
	}

	user := userData.ToUser()

	return &user, nil
}

//FindUserByEmail If data not found will return nil without error
func (repo *GormRepository) FindUserByEmail(email string) (*user.User, error) {

	var userData User

	err := repo.DB.Where("email = ?", email).First(&userData).Error
	if err != nil {
		return nil, err
	}

	user := userData.ToUser()

	return &user, nil
}

//FindAllUser find all user with given specific page and row per page, will return empty slice instead of nil
func (repo *GormRepository) FindAllUser(skip int, rowPerPage int) ([]user.User, error) {

	var users []User

	err := repo.DB.Offset(skip).Limit(rowPerPage).Find(&users).Error
	if err != nil {
		return nil, err
	}

	var result []user.User
	for _, value := range users {
		result = append(result, value.ToUser())
	}

	return result, nil
}

//InsertUser Insert new User into storage
func (repo *GormRepository) InsertUser(user user.User) error {

	userData := newUserTable(user)
	userData.ID = 0

	err := repo.DB.Create(userData).Error
	if err != nil {
		return err
	}

	return nil
}

//UpdateItem Update existing item in database
func (repo *GormRepository) UpdateUser(user user.User, currentVersion int) error {

	userData := newUserTable(user)

	err := repo.DB.Model(&userData).Updates(User{Name: userData.Name}).Error
	if err != nil {
		return err
	}

	return nil
}
