package database

import (
	"log"
	"logify/configs"

	"gorm.io/gorm"
)

func Database() *gorm.DB {
	db, err := configs.ConnectDb()
	if err != nil {
		log.Fatal(err)
	}

	return db
}
