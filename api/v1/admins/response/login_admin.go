package response

import (
	"altaStore/business/admins"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

type AdminLogin struct {
	ID           int
	Name         string
	Status       string
	Email        string
	Phone_number string
	Username     string
	CreatedBy    string
	Token        string
	Created_at   time.Time
	Updated_at   time.Time
	Deleted_at   time.Time
}

func ConvertAdminToAdminLogin(admin *admins.Admins) *AdminLogin {
	token, _ := createToken(admin)
	return &AdminLogin{
		ID:           admin.ID,
		Name:         admin.Name,
		Status:       admin.Status,
		Email:        admin.Email,
		Phone_number: admin.Phone_number,
		Username:     admin.Username,
		CreatedBy:    admin.CreatedBy,
		Token:        token,
		Created_at:   admin.Created_at,
		Updated_at:   admin.Updated_at,
		Deleted_at:   admin.Deleted_at,
	}
}

func createToken(admin *admins.Admins) (string, error) {
	claims := jwt.MapClaims{
		"id":         admin.ID,
		"created_at": admin.Created_at,
		"authorized": true,
		"timestamp":  time.Now(),
		"exp":        time.Now().Add(time.Hour * 1).Unix(),
		"name":       admin.Name,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}
