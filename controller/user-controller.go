package controller

import (
	"go-api/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserController interface {
	GetUser(ctx *gin.Context) error
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
