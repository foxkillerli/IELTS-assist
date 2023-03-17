package db

import (
	"fmt"
	"github.com/foxkillerli/IELTS-assist/config"
	"github.com/jinzhu/gorm"
	"log"
)

var DB *gorm.DB

func GetInstance() *gorm.DB {
	if DB == nil {
		ConnectionString := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=%s", config.PostgresHost, config.PostgresPort, config.PostgresUser, config.PostgresDBName, config.PostgresPassword, "disable")
		DB, _ = gorm.Open("postgres", ConnectionString)
		DB.DB().SetMaxOpenConns(100)
		DB.DB().SetMaxIdleConns(20)
	}
	return DB
}

func init() {
	var err error
	ConnectionString := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=%s", config.PostgresHost, config.PostgresPort, config.PostgresUser, config.PostgresDBName, config.PostgresPassword, "disable")
	DB, err := gorm.Open("postgres", ConnectionString)
	if err != nil {
		log.Printf("postgres connect error %v", err)
	}
	if DB.Error != nil {
		log.Printf("database error %v", DB.Error)
	}
}
