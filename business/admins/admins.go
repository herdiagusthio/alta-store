package admins

import "time"

// Admins data core
type Admins struct {
	ID     int
	Name   string
	Status string
	// Password     string
	Email        string
	Phone_number string
	Username     string
	CreatedBy    string
	Modified_by  string
	Created_at   time.Time
	Updated_at   time.Time
	Password     string
	Deleted_at   time.Time
}

func NewAdmin(
	name string,
	username string,
	email string,
	password string,
	phone_number string,
	created_by string,
) Admins {

	return Admins{
		Name:         name,
		Username:     username,
		Email:        email,
		Phone_number: phone_number,
		Created_at:   time.Now(),
		CreatedBy:    created_by,
		Status:       "active",
		Password:     password,
	}
}

func (old_admin *Admins) ModifyAdmin(new_data AdminUpdatable) Admins {
	return Admins{
		ID:           old_admin.ID,
		Status:       new_data.Status,
		Name:         new_data.Name,
		Email:        old_admin.Email,
		Phone_number: new_data.Phone_number,
		Username:     old_admin.Username,
		CreatedBy:    old_admin.CreatedBy,
		Modified_by:  new_data.Modified_by,
		Created_at:   old_admin.Created_at,
		Updated_at:   time.Now(),
	}
}
