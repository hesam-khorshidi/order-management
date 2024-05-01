package providers

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"order-management/common"
)

func ConnectToDB() *gorm.DB {

	host := common.GetEnv("DB_HOST", "localhost")
	dbPort := common.GetEnv("DB_PORT", "5432")
	dbUsername := common.GetEnv("DB_USER", "")
	dbName := common.GetEnv("DB_NAME", "")
	dbPassword := common.GetEnv("DB_PASSWORD", "")

	// Database connection string
	dsn := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s port=%s", host, dbUsername, dbName, dbPassword, dbPort)

	// Opening connection to database
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Panicf("error connecting to database: %s", err.Error())
	}

	return db
}
