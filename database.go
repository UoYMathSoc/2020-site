package main

import (
	"fmt"
	"log"

	"github.com/UoYMathSoc/2020-site/models"
	"github.com/UoYMathSoc/2020-site/structs"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Models contains all active models used the database
var Models = [...]interface{}{
	models.UserModel{},
	models.EventModel{},
}

// InitDatabase creates a connection to the database based on the config c.
func InitDatabase(c *structs.Config) *gorm.DB {
	dsn := fmt.Sprintf("%s", c.Server.Database)
	db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Print(err)
		panic("failed to connect to database")
	}
	log.Printf("Connection opened to Database: %s", c.Server.Database)
	for _, model := range Models {
		err = db.AutoMigrate(model)
		if err != nil {
			log.Fatalf("could not migrate database: %s", err)
		}
	}
	log.Printf("Database Migrated")
	return db
}
