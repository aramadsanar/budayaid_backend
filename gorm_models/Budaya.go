package gorm_models

type Budaya struct {
	Id               int    `gorm:"primary_key" db:"id" json:"id"`
	Name             string `db:"name" json:"name" gorm:"type:text"`
	Description      string `db:"description" json:"description" gorm:"type:text"`
	ImageUrl         string `db:"image_url" json:"image_url" gorm:"type:text"`
	GoogleSearchTerm string `db:"google_search_term" json:"google_search_term" gorm:"type:text"`

	Province   Province   `gorm:"foreignkey:ProvinceID;association_foreignkey:id" json:"-"`
	ProvinceID int        `json:"province_id"`
	Categories Categories `gorm:"foreignkey:Category;association_foreignkey:id" json:"-"`
	Category   int        `json:"category"`
}

func (Budaya) TableName() string {
	return "budaya"
}
