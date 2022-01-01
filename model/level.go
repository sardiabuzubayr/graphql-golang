package model

import (
	"mygraphql/config"

	"gorm.io/gorm"
)

type LevelUser struct {
	IdLevel int    `json:"id_level"`
	Nama    string `json:"nama"`
}

func (LevelUser) TableName() string {
	return "level_user"
}

func GetLevel(idLevel int) ([]LevelUser, error) {
	levels := new([]LevelUser)
	var result *gorm.DB
	if idLevel != 0 {
		result = config.DB.Where("id_level = ? ", idLevel).Find(levels)
	} else {
		result = config.DB.Find(levels)
	}
	return *levels, result.Error
}
