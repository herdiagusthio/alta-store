package migration

import (
	"altaStore/modules/address"
	"altaStore/modules/cart"
	"altaStore/modules/user"

	"gorm.io/gorm"
)

func InitMigrate(db *gorm.DB) {
	db.AutoMigrate(
		&user.User{},
		&address.Address{},
		&cart.Cart{},
		&cart.CartDetail{},
	)
}
