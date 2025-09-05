package service

import (
	"go-api/entity"

	"gorm.io/gorm"
)

type AuthService interface {
	Register(user entity.RegisterRequest) (entity.RegisterRequest, error)
	Login(email entity.LoginRequest) (entity.User, error)
}

type authService struct {
	db *gorm.DB
}

func NewAuthService(db *gorm.DB) AuthService {
	return &authService{
		db: db,
	}
}

func (ctx authService) Register(user entity.RegisterRequest) (entity.RegisterRequest, error) {
	var userEntity entity.User = entity.User{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Age:       user.Age,
		Email:     user.Email,
		Password:  user.Password,
	}
	err := ctx.db.Create(&userEntity).Error

	return user, err
}

func (ctx authService) Login(request entity.LoginRequest) (entity.User, error) {
	var user entity.User

	err := ctx.db.Where("email = ?", request.Email).First(&user).Error

	return user, err
}
