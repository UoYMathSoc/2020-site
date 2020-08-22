package database

import (
	"github.com/UoYMathSoc/2020-site/structs"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"log"
)

var Instance *gorm.DB

func Init(config *structs.Config, models ...interface{}) {
	var err error
	Instance, err = gorm.Open("sqlite3", config.Server.Database)
	if err != nil {
		log.Print(err)
		panic("failed to connect to database")
	}
	log.Printf("Connection opened to Database: %s", config.Server.Database)
	for i, model := range models {
		Instance.AutoMigrate(model)
		var percent int = (i * 100) / len(models)
		log.Printf("Database Migrated: %d%", percent)
	}
	log.Printf("Database Migrated")
}
