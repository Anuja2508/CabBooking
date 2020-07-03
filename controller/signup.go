package controller

import (
	"GoCab/model"

	"github.com/jinzhu/gorm"
)

type User struct {
	database *gorm.DB
}

// NewUserController will create new user controller
func NewUserController(database *gorm.DB) User {
	return User{
		database: database,
	}
}

// Check :- For checking if phone number and email exit in database
func (controller User) Check(Email string, Phone string) (bool, bool) {

	emailCheck := controller.database.Where("email ilike ?", Email).First(&model.User{}).RecordNotFound()
	phoneCheck := controller.database.Where("phone = ?", Phone).First(&model.User{}).RecordNotFound()
	return emailCheck, phoneCheck

}

// Create For creating User in database
func (controller User) Create(user *model.User) error {

	if err := controller.database.Create(user).Error; err != nil {
		return err
	}
	return nil

}

// FindUser for finding user from database
func (controller User) FindUser(Phone string) (*model.User, error) {
	user := new(model.User)

	if err := controller.database.Where("mobile_number = ?", Phone).Find(&model.User{}).Scan(user).Error; err != nil {
		return user, err
	}
	return user, nil

}

// RegisterCab For registering cab
func (controller User) RegisterCab(cab *model.Cab) error {

	if err := controller.database.Create(cab).Error; err != nil {
		return err
	}
	return nil

}
