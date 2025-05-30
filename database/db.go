package db

import (
	"fmt"
	"log"

	"github.com/Chenorlive/brainy/model"

	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func NewSqliteGorm() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("core.db"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	return db, nil
}

func NewMysqlGorm() (*gorm.DB, error) {
	dsn := "root:root@tcp(localhost:3306)/ecom?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	return db, nil
}

func SetupDB(debug bool) (*gorm.DB, error) {

	var db *gorm.DB
	var err error
	if debug {
		db, err = NewSqliteGorm()
	} else {
		db, err = NewSqliteGorm()
	}

	if err != nil {
		log.Fatal(err)
	}

	// all model

	// Migrate the schema

	for _, model := range model.Models {
		if err := db.AutoMigrate(model); err != nil {
			log.Fatal("Failed to auto-migrate table:", err)
		}
		fmt.Printf("Auto-migrated table for model: %T\n", model)
	}

	if err != nil {
		log.Fatal(err)
	}

	return db, nil
}
