package service

import (
	"errors"
	"fmt"
	"go-api/entity"

	"gorm.io/gorm"
)

type UserService interface {
	GetUser(id string) (entity.User, error)
	UpdateUser(id string, req entity.UpdateUserRequest) (entity.User, error)
	GetAllUsers() ([]entity.UserResponse, error)
	DeleteUser(id string) error
}

type userService struct {
	db *gorm.DB
}

func NewUserService(db *gorm.DB) UserService {
	return &userService{
		db: db,
	}
}

func (ctx userService) GetUser(id string) (entity.User, error) {

	var user entity.User
	err := ctx.db.Where("id = ?", id).First(&user).Error

	return user, err
}

func (ctx userService) GetAllUsers() ([]entity.UserResponse, error) {

	var users []entity.User
	err := ctx.db.Find(&users).Error

	var response []entity.UserResponse
	for _, u := range users {
		response = append(response, entity.UserResponse{
			ID:        u.ID,
			FirstName: u.FirstName,
			LastName:  u.LastName,
			Age:       u.Age,
			CreatedAt: u.CreatedAt,
			UpdatedAt: u.UpdatedAt,
			Email:     u.Email,
		})
	}

	return response, err
}

func (ctx userService) UpdateUser(id string, req entity.UpdateUserRequest) (entity.User, error) {

	var user entity.User
	err := ctx.db.First(&user, "id = ?", id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return entity.User{}, fmt.Errorf("user with ID %s not found", id)
	}

	if err != nil {
		return entity.User{}, fmt.Errorf("database error: %v", err)
	}
	update := map[string]interface{}{}

	if req.FirstName != "" {
		update["first_name"] = req.FirstName
	}
	if req.LastName != "" {
		update["last_name"] = req.LastName
	}
	if req.Age > 0 && req.Age <= 120 {
		update["age"] = req.Age
	}

	if len(update) == 0 {
		return user, nil // No fields to update
	}

	err = ctx.db.Model(&user).Updates(update).Error
	if err != nil {
		return entity.User{}, fmt.Errorf("failed to update user: %v", err)
	}

	return user, err
}

func (ctx userService) DeleteUser(id string) error {

	result := ctx.db.Delete(&entity.User{}, "id = ?", id)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}
