package actions

import (
	"budayaid/budayaid_backend/gorm_models"
	"github.com/gobuffalo/buffalo"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"strconv"
)

// HomeHandler is a default handler to serve up
// a home page.
func HomeHandler(c buffalo.Context) error {
	return c.Render(200, r.HTML("index.html"))
}

func InitDB(c buffalo.Context) error {
	error_message := gorm_models.InitMigration()
	return c.Render(200, r.String(error_message))
}

func PukeBudaya(c buffalo.Context) error {
	db, _ := gorm.Open("sqlite3", "budayaid.db")
	defer db.Close()

	budayas := []gorm_models.Budaya{}
	db.Find(&budayas)

	return c.Render(200, r.JSON(budayas))
}

func GetBudayasByProvinceId(c buffalo.Context) error {
	db, _ := gorm.Open("sqlite3", "budayaid.db")
	defer db.Close()

	result := []gorm_models.Budaya{}

	province_id, _ := strconv.Atoi(c.Param("province_id"))
	db.Where(&gorm_models.Budaya{ProvinceID: province_id}).Find(&result)

	return c.Render(200, r.JSON(result))
}

func GetBudayasByProvinceIdAndCategory(c buffalo.Context) error {
	db, _ := gorm.Open("sqlite3", "budayaid.db")
	defer db.Close()

	result := []gorm_models.Budaya{}

	province_id, _ := strconv.Atoi(c.Param("province_id"))
	category_id, _ := strconv.Atoi(c.Param("category_id"))
	db.Where(&gorm_models.Budaya{ProvinceID: province_id, Category: category_id}).Find(&result)

	return c.Render(200, r.JSON(result))
}

func GetProvinces(c buffalo.Context) error {
	db, _ := gorm.Open("sqlite3", "budayaid.db")
	defer db.Close()

	result := []gorm_models.Province{}

	db.Find(&result)
	return c.Render(200, r.JSON(result))
}

///getJSONProvincesByIslandName/<string:islandName>
func GetProvincesByIslandName(c buffalo.Context) error {
	db, _ := gorm.Open("sqlite3", "budayaid.db")
	defer db.Close()

	result := []gorm_models.Province{}

	db.Where(gorm_models.Province{Island: c.Param("island_name")}).Find(&result)
	return c.Render(200, r.JSON(result))
}
