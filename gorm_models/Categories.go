package gorm_models

type Categories struct {
	Id           int    `gorm:"primary_key;type:text" db:"id" json:"id"`
	Name         string `db:"name" json:"name" gorm:"type:text"`
	FriendlyName string `db:"friendly_name" json:"friendly_name" gorm:"type:text"`
}

func (Categories) TableName() string {
	return "categories"
}
