package database

import (
	"github.com/Kulaxyz/partner/config"
	"gorm.io/driver/postgres"

	//"api-fiber-gorm/model"
	"fmt"
	_ "gorm.io/driver/postgres"
	"gorm.io/gorm"
	"strconv"
)

// ConnectDB connect to db
func ConnectDB() {
	p := config.Config("DB_PORT")
	port, err := strconv.ParseUint(p, 10, 32)

	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", config.Config("DB_HOST"), port, config.Config("DB_USER"), config.Config("DB_PASSWORD"),config.Config("DB_NAME"))
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database" + err.Error())
	}

	fmt.Println("Connection Opened to Database")
}

