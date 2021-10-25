package migration

import (
	"altaStore/modules/admins"

	"gorm.io/gorm"
)

func InitMigrate(db *gorm.DB) {
	db.AutoMigrate(&admins.AdminsTable{})
}
