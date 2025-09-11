package main

import (
	"go-api/controller"
	"go-api/database"
	"go-api/middlewares"
	"go-api/service"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	server := gin.New()

	server.SetTrustedProxies([]string{})

	DB = database.ConnectDatabase()

	if DB == nil {
		panic("Failed to connect to database")
	}

	var (
		videoService    service.VideoService       = service.NewVideoService(DB)
		authService     service.AuthService        = service.NewAuthService(DB)
		userService     service.UserService        = service.NewUserService(DB)
		videoController controller.VideoController = controller.NewVideoController(videoService)
		userController  controller.UserController  = controller.NewUserController(userService)
		authController  controller.AuthController  = controller.NewAuthController(authService)
	)

	server.Use(gin.Recovery(), middlewares.Logger())

	server.POST("/api/signup", func(ctx *gin.Context) {
		authController.Register(ctx)
	})

	server.POST("/api/login", func(ctx *gin.Context) {
		authController.Login(ctx)
	})

	// Auth Middleware applies to api group alone
	apiRoutes := server.Group("/api", middlewares.AuthMiddleware())
	{
		apiRoutes.GET("/videos", func(ctx *gin.Context) {
			videoController.GetAll(ctx)
		})
		apiRoutes.POST("/videos", func(ctx *gin.Context) {
			videoController.SaveVideo(ctx)
		})

		apiRoutes.DELETE("/videos/:id", func(ctx *gin.Context) {
			videoController.DeleteVideo(ctx)
		})
		apiRoutes.GET("/videos/:id", func(ctx *gin.Context) {
			videoController.GetVideo(ctx)
		})
		apiRoutes.PUT("/videos/:id", func(ctx *gin.Context) {
			videoController.UpdateVideo(ctx)
		})
		apiRoutes.GET("/user/:id", func(ctx *gin.Context) {
			userController.GetUser(ctx)
		})
		apiRoutes.GET("/users", func(ctx *gin.Context) {
			userController.GetAllUsers(ctx)
		})
		apiRoutes.PUT("/user/:id", func(ctx *gin.Context) {
			userController.UpdateUser(ctx)
		})
		apiRoutes.DELETE("/user/:id", func(ctx *gin.Context) {
			userController.DeleteUser(ctx)
		})
	}

	server.Run(":8080")
}
