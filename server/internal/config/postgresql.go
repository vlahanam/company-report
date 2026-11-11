package config

import (
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func SetupDB(dbconfig *DbConfig) *gorm.DB {
	db, err := gorm.Open(postgres.Open(dbconfig.dsn()), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Failed to get sql.DB: %v", err)
	}

	// Maximum number of idle connections
	sqlDB.SetMaxIdleConns(25)
	// Maximum number of open connections
	sqlDB.SetMaxOpenConns(50)
	// Maximum amount of time a connection may be reused
	sqlDB.SetConnMaxLifetime(time.Hour)

	return db
}
