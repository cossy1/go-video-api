package controller

import (
	"go-api/entity"
	"go-api/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type VideoController interface {
	SaveVideo(ctx *gin.Context) error
	GetAll(ctx *gin.Context) error
	GetVideo(ctx *gin.Context) error
	UpdateVideo(ctx *gin.Context) error
}

type videoController struct {
	service service.VideoService
}

func NewVideoController(service service.VideoService) VideoController {
	// if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
	// 	v.RegisterValidation("is-cool", validators.ValidateCoolTitle)
	// }

	return &videoController{service: service}
}

func (vc *videoController) GetAll(ctx *gin.Context) error {

	id, exists := ctx.Get("userId")

	if !exists {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Unauthorized"})
		return nil
	}

	userId := id.(string)

	data, err := vc.service.GetAll(userId)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

		return err
	}

	var response []entity.VideoResponse
	for _, u := range data {
		response = append(response, *entity.ToVideoResponse(u))
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Videos fetched successfully!", "data": response})

	return nil
}

func (vc *videoController) SaveVideo(ctx *gin.Context) error {
	var video entity.CreateVideoRequest

	err := ctx.ShouldBindJSON(&video)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return err
	}

	userId, exists := ctx.Get("userId")

	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return nil
	}

	parsedUUID, err := uuid.Parse(userId.(string))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return err
	}
	video.UserID = parsedUUID

	savedVideo, err := vc.service.SaveVideo(video)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

		return err
	}

	response := entity.ToVideoResponse(savedVideo)

	ctx.JSON(http.StatusCreated, gin.H{"message": "Video saved successfully", "data": response})

	return nil

}

func (vc *videoController) GetVideo(ctx *gin.Context) error {

	id := ctx.Param("id")

	if id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Id is required"})
		return nil
	}

	data, err := vc.service.GetVideo(id)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

		return err
	}

	response := entity.ToVideoResponse(data)

	ctx.JSON(http.StatusOK, gin.H{"message": "Video fetched successfully!", "data": response})

	return nil
}

func (vc *videoController) UpdateVideo(ctx *gin.Context) error {
	var req entity.UpdateVideoRequest

	err := ctx.ShouldBindJSON(&req)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return err
	}

	id := ctx.Param("id")

	if id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Id is required"})
		return nil
	}

	data, err := vc.service.UpdateVideo(id, req)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

		return err
	}

	response := entity.ToVideoResponse(data)

	ctx.JSON(http.StatusOK, gin.H{"message": "Video updated successfully!", "data": response})

	return nil
}
