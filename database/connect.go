package database

import (
	"fmt"
	"log"
	"technicaltest/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func ConnectDB() {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", config.Conf.DatabaseHost, config.Conf.DatabasePort, config.Conf.DatabaseUser, config.Conf.DatabasePass, config.Conf.DatabaseName)
	log.Println(dsn)
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true,
	}), &gorm.Config{
		PrepareStmt: true,
		Logger:      logger.Default.LogMode(logger.Silent),
	})

	if err != nil {
		log.Fatal(err)
	}

	DB = db
}
