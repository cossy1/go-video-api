package service

import (
	"go-api/entity"

	"gorm.io/gorm"
)

type UserService interface {
	GetUser(id uint64) (entity.User, error)
	// UpdateUser(id uint64) (entity.User, error)
	// GetAllUsers() ([]entity.User, error)
	// DeleteUser(id uint64) error
}

type userService struct {
	db *gorm.DB
}

func NewUserService(db *gorm.DB) UserService {
	return &userService{
		db: db,
	}
}

func (ctx userService) GetUser(id uint64) (entity.User, error) {

	var user entity.User
	err := ctx.db.Where("ID = ?", id).First(&user).Error

	return user, err
}

// func GetAllUsers() (entity.User, error) {

// 	var users entity.User
// 	err := database.DB.Find(&users).Error

// 	return users, err
// }

// func UpdateUser(id uint64, user entity.User) (entity.User, error) {

// 	var users entity.User
// 	err := database.DB.Update(id, user).Error

// 	return users, err
// }

// func DeleteUser(id uint64) error {

// 	var users entity.User
// 	err := database.DB.Delete(&users.ID, id).Error

// 	return err
// }
