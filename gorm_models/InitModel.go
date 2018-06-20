package gorm_models

import (
	_ "fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func InitMigration() string {
	db, err := gorm.Open("sqlite3", "budayaid.db")
	defer db.Close()

	if err != nil {
		return err.Error()
	}

	db.AutoMigrate(&Categories{}, &Province{}, &Budaya{})
	return ""
}
