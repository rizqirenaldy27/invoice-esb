package database

import (
	"fmt"

	"github.com/rizqirenaldy27/invoice-esb/infrastructure/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB() (*gorm.DB, error) {
	dsn := config.AppConfig.DatabaseConfig.DatabaseURI

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	if err := sqlDB.Ping(); err != nil {
		return nil, err
	}

	fmt.Println("Database Connected")

	return db, nil
}
