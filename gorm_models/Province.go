package gorm_models

import ()

type Province struct {
	Id           int    `gorm:"primary_key" db:"id" json:"id"`
	Name         string `db:"name" json:"name" gorm:"type:varchar(250)"`
	FriendlyName string `db:"friendly_name" json:"firendly_name" gorm:"type:varchar(250)"`
	Island       string `db:"island" json:"island" gorm:"type:varchar(50)"`
}

func (Province) TableName() string {
	return "province"
}
