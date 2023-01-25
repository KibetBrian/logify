package configs

import (
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"logify/models"
)

type Configurations struct {
	DB_HOST     string
	DB_PORT     string
	DB_USER     string
	DB_PASSWORD string
	DB_NAME     string
}

var configurations Configurations

func init() {

	viper.SetConfigName("config")
	viper.AddConfigPath("../")
	viper.AutomaticEnv()
	viper.SetEnvPrefix("spruce")

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %s ", err))
	}

	err = viper.Unmarshal(&configurations)
	if err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
	}

}

var Db *gorm.DB

func ConnectDb() (*gorm.DB, error) {

	//Postgres connection string
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", configurations.DB_HOST, configurations.DB_PORT, configurations.DB_USER, configurations.DB_USER, configurations.DB_NAME)

	db, err := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("database connection failed: %v", err)
	}

	//Auto migrations
	db.AutoMigrate(&models.Product{})
	return db, nil
}
