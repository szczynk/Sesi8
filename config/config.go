package config

import (
	"fmt"

	"github.com/szczynk/Sesi8/structs"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func DBInit() *gorm.DB {
	var (
		dbHost = "localhost"
		dbPort = "5432"
		dbUser = "postgres"
		dbPass = "admin"
		dbName = "living-code"
		db     *gorm.DB
		err    error
	)
	config := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta", dbHost, dbUser, dbPass, dbName, dbPort)
	db, err = gorm.Open(postgres.Open(config), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}

	db.AutoMigrate(structs.Person{})
	return db
}
