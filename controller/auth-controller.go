package controller

import (
	"go-api/entity"
	"go-api/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type AuthController interface {
	Register(ctx *gin.Context) error
	Login(ctx *gin.Context) error
}

type authController struct {
	service service.AuthService
}

func NewAuthController(service service.AuthService) AuthController {
	return &authController{service: service}
}

func (ac *authController) Register(ctx *gin.Context) error {
	var request entity.RegisterRequest

	err := ctx.ShouldBindJSON(&request)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return err
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), 10)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return err
	}

	request.Password = string(hashedPassword)

	data, err := ac.service.Register(request)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

		return err
	}

	response := map[string]interface{}{
		"id":        data.ID,
		"age":       data.Age,
		"firstName": data.FirstName,
		"lastName":  data.LastName,
		"email":     data.Email,
		"createdAt": data.CreatedAt,
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "User created successfully", "data": response})

	return nil

}

func (ac *authController) Login(ctx *gin.Context) error {
	var payload entity.LoginRequest

	err := ctx.ShouldBindJSON(&payload)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return err
	}

	data, err := ac.service.Login(payload)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "User not found"})
		return err
	}

	passwordMatches := bcrypt.CompareHashAndPassword([]byte(data.Password), []byte(payload.Password))

	if passwordMatches != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})

		return nil
	}

	accessToken, err := service.GenerateToken(data.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return err
	}

	response := map[string]interface{}{
		"access_token": accessToken,
		"user":         entity.ToUserResponse(data),
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Login successful!", "data": response})

	return nil
}
