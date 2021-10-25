package response

import "altaStore/business/admins"

type DataAdmin struct {
	Name string
	// Status       string
	Email        string
	Phone_number string
	Username     string
}

func ConvertAdminToDataAdmin(admin *admins.Admins) *DataAdmin {
	return &DataAdmin{
		Name: admin.Name,
		// Status:       admin.Status,
		Username:     admin.Username,
		Email:        admin.Email,
		Phone_number: admin.Phone_number,
	}
}

func ConvertListAdminToDataAdmin(admins *[]admins.Admins) []DataAdmin {
	var list_data_admin []DataAdmin
	for _, data := range *admins {
		list_data_admin = append(list_data_admin, *ConvertAdminToDataAdmin(&data))
	}
	return list_data_admin
}
