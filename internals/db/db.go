package db

import (
	"OfflineSearchEngine/configs"
	"OfflineSearchEngine/models"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func LoadDb(conf *configs.DbConfigs) {
	db, err := gorm.Open(postgres.Open(fmt.Sprintf("%s://%s:%s@%s:%d/%s",
		conf.Driver, conf.User, conf.Password, conf.IP, conf.Port, conf.Name)), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&models.User{})

}
