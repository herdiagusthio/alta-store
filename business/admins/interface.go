package admins

type Service interface {
	FindAllAdmin(offset, limit int) (*[]Admins, error)
	FindAdminByUsername(username string) (*Admins, error)
	FindAdminById(id_admins int) (*Admins, error)
	LoginAdmin(username, password string) (*Admins, error)
	InsertAdmin(admin_spec AdminSpec, createdBy string) error
	ModifyAdmin(username string, admin AdminUpdatable) error
}

type Repository interface {
	GetAdmin(limit, offset int) (*[]Admins, error)
	GetAdminByUsername(username string) (*Admins, error)
	GetAdminById(id int) (*Admins, error)
	LoginAdmin(username, password string) (*Admins, error)
	CreateAdmin(admin *Admins) error
	UpdateAdmin(admin *Admins) error
}
