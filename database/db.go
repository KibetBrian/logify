package database

import (
	"log"
	"shoppy/configs"

	"gorm.io/gorm"
)

func Database() *gorm.DB {
	db, err := configs.ConnectDb()
	if err != nil {
		log.Fatal(err)
	}

	return db
}
