package migration

import (
	"gorm.io/gorm"
)

func InitMigrate(db *gorm.DB) {
	// db.AutoMigrate(&user.UserTable{}, &pet.PetTable{})
}
