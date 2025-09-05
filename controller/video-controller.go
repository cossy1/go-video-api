package controller

import (
	"go-api/entity"
	"go-api/service"
	"go-api/validators"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

type VideoController interface {
	SaveVideo(ctx *gin.Context) error
	FindAll() []entity.Video
}

type videoController struct {
	service service.VideoService
}

func NewVideoController(service service.VideoService) VideoController {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("is-cool", validators.ValidateCoolTitle)
	}

	return &videoController{service: service}
}

func (vc *videoController) FindAll() []entity.Video {
	return vc.service.FindAll()
}

func (vc *videoController) SaveVideo(ctx *gin.Context) error {
	var video entity.Video

	err := ctx.ShouldBindJSON(&video)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return err
	}

	// err = validate.Struct(video)

	// if err != nil {
	// 	ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

	// 	return err
	// }

	savedVideo := vc.service.SaveVideo(video)

	ctx.JSON(http.StatusCreated, gin.H{"message": "Video saved successfully", "data": savedVideo})

	return nil

}
