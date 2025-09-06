package controller

import (
	"go-api/entity"
	"go-api/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserController interface {
	GetUser(ctx *gin.Context) error
	GetAllUsers(ctx *gin.Context) error
	UpdateUser(ctx *gin.Context) error
	DeleteUser(ctx *gin.Context) error
}

type userController struct {
	service service.UserService
}

func NewUserController(service service.UserService) UserController {
	return &userController{service: service}
}

func (uc *userController) GetUser(ctx *gin.Context) error {

	id := ctx.Param("id")

	if id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "User Id is required"})
		return nil
	}
	// Convert id string to uint64
	userId, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return err
	}

	data, err := uc.service.GetUser(userId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "User not found"})

		return err
	}

	response := map[string]interface{}{
		"id":        data.ID,
		"age":       data.Age,
		"firstName": data.FirstName,
		"lastName":  data.LastName,
		"email":     data.Email,
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "User fetched successfully", "data": response})

	return nil

}

func (uc *userController) GetAllUsers(ctx *gin.Context) error {

	data, err := uc.service.GetAllUsers()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "An error occured"})

		return err
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Users fetched successfully", "data": data})

	return nil

}

func (uc *userController) UpdateUser(ctx *gin.Context) error {
	var body entity.UpdateUserRequest
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return err
	}
	idx := ctx.Param("id")

	id, err := strconv.ParseUint(idx, 10, 64)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return err
	}

	data, err := uc.service.UpdateUser(id, body)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "An error occured"})

		return err
	}

	response := map[string]interface{}{
		"id":        data.ID,
		"age":       data.Age,
		"firstName": data.FirstName,
		"lastName":  data.LastName,
		"email":     data.Email,
		"updatedAt": data.UpdatedAt,
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "User updated successfully", "data": response})

	return nil

}

func (uc *userController) DeleteUser(ctx *gin.Context) error {

	id := ctx.Param("id")

	if id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "User Id is required"})
		return nil
	}

	userId, err := strconv.ParseUint(id, 10, 64)

	if err != nil {

		return err
	}

	if err := uc.service.DeleteUser(userId); err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return err
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})

	return nil

}
