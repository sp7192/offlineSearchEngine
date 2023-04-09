package db

import (
	"OfflineSearchEngine/configs"
	"OfflineSearchEngine/models"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DatabaseHandler struct {
	db *gorm.DB
}

func LoadDb(conf *configs.DbConfigs) *DatabaseHandler {
	db, err := gorm.Open(postgres.Open(fmt.Sprintf("%s://%s:%s@%s:%d/%s",
		conf.Driver, conf.User, conf.Password, conf.IP, conf.Port, conf.Name)), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	// Migrate the schema
	db.AutoMigrate(&models.User{})

	return &DatabaseHandler{
		db: db,
	}
}

func (handler *DatabaseHandler) CreateUser(user models.User) error {
	return nil
}

func (handler *DatabaseHandler) RemoveUser(user models.User) error {
	return nil
}

func (handler *DatabaseHandler) FindUserName(user models.User) bool {
	return true
}

func (handler *DatabaseHandler) UpdateUser(user models.User) error {
	return nil
}
