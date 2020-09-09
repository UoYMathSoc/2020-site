package main

import (
	"log"

	"github.com/UoYMathSoc/2020-site/models"
	"github.com/UoYMathSoc/2020-site/structs"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// Models contains all active models used the database
var Models = [...]interface{}{
	models.UserModel{},
}

// InitDatabase creates a connection to the database based on the config c.
func InitDatabase(c *structs.Config) *gorm.DB {
	db, err := gorm.Open("sqlite3", c.Server.Database)
	if err != nil {
		log.Print(err)
		panic("failed to connect to database")
	}
	log.Printf("Connection opened to Database: %s", c.Server.Database)
	for i, model := range Models {
		db.AutoMigrate(model)
		var percent int = (i * 100) / len(Models)
		log.Printf("Database Migrated: %d%%", percent)
	}
	log.Printf("Database Migrated!")
	return db
}
