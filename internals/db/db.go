package db

import (
	"OfflineSearchEngine/configs"
	"OfflineSearchEngine/internals/utils"
	"OfflineSearchEngine/models"
	"errors"
	"fmt"
	"log"

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
		panic("failed to connect database, create the db first!")
	}
	// Migrate the schema
	db.AutoMigrate(&models.User{})
	return &DatabaseHandler{
		db: db,
	}
}

func (handler *DatabaseHandler) InitUsers() {
	hashedPass, err := utils.HashAndSalt([]byte("1234"))
	if err != nil {
		log.Fatal("Could not hash and salt password")
	}

	handler.CreateUser(&models.User{Username: "admin", Password: hashedPass})
}

func (handler *DatabaseHandler) CreateUser(user *models.User) (*models.User, error) {
	res := handler.db.Create(user)
	if res.RowsAffected == 0 {
		return &models.User{}, errors.New("user not found")
	}
	return user, nil
}

func (handler *DatabaseHandler) ReadUser(id string) (*models.User, error) {
	var user models.User
	res := handler.db.First(&user, id)
	if res.RowsAffected == 0 {
		return nil, errors.New("article not found")
	}
	return &user, nil
}

func (handler *DatabaseHandler) FindUser(user *models.User) bool {
	res := handler.db.Find(&user)
	if res.RowsAffected == 0 {
		return false
	}
	return true
}

func (handler *DatabaseHandler) ReadUsers() ([]*models.User, error) {
	var users []*models.User
	res := handler.db.Find(&users)
	if res.Error != nil {
		return nil, errors.New("authors not found")
	}
	return users, nil
}

func (handler *DatabaseHandler) UpdateUser(user *models.User) (*models.User, error) {

	var updateUser models.User
	result := handler.db.Model(&updateUser).Where(user.ID).Updates(user)
	if result.RowsAffected == 0 {
		return &models.User{}, errors.New("artcile not updated")
	}
	return &updateUser, nil
}

func (handler *DatabaseHandler) DeleteUser(id string) error {
	var deleteUser models.User
	result := handler.db.Where(id).Delete(&deleteUser)
	if result.RowsAffected == 0 {
		return errors.New("article data not deleted")
	}
	return nil
}
