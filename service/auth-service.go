package service

import (
	"go-api/entity"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AuthService interface {
	Register(user entity.RegisterRequest) (entity.User, error)
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

func (ctx authService) Register(user entity.RegisterRequest) (entity.User, error) {
	userEntity := entity.User{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Age:       user.Age,
		Email:     user.Email,
		Password:  user.Password,
	}
	err := ctx.db.Create(&userEntity).Error

	return userEntity, err
}

func (ctx authService) Login(request entity.LoginRequest) (entity.User, error) {
	var user entity.User

	err := ctx.db.Where("email = ?", request.Email).First(&user).Error

	return user, err
}

func GenerateUUID() uuid.UUID {
	return uuid.New()
}
