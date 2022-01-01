package model

import "mygraphql/config"

type Users struct {
	Email       string `json:"email" form:"email" gorm:"primaryKey"`
	Nama        string `json:"nama" form:"nama"`
	NoHandphone string `json:"no_handphone" form:"no_handphone"`
	Alamat      string `json:"alamat" form:"alamat"`
	Ktp         string `json:"ktp" form:"ktp"`
}

func GetUsers(keywords string) ([]Users, error) {
	var users []Users
	result := config.DB.Where("email LIKE ? OR nama LIKE ?", "%"+keywords+"%", "%"+keywords+"%").Find(&users)

	return users, result.Error
}
