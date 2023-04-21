package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/mygram/MyGram-API/controllers"
	"github.com/mygram/MyGram-API/middlewares"
	"github.com/mygram/MyGram-API/models"
	"github.com/mygram/MyGram-API/repositories"
	"github.com/mygram/MyGram-API/utils"
)

func main() {
	// Initialize Gin router
	router := gin.Default()

	// Connect to database
	db := utils.ConnectDB()
	defer db.Close()

	// Migrate models to database
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Photo{})
	db.AutoMigrate(&models.Comment{})
	db.AutoMigrate(&models.SocialMedia{})

	// Setup repositories
	userRepo := repositories.NewUserRepository(db)
	photoRepo := repositories.NewPhotoRepository(db)
	commentRepo := repositories.NewCommentRepository(db)
	socialMediaRepo := repositories.NewSocialMediaRepository(db)

	// Setup controllers
	authController := controllers.NewAuthController(userRepo)
	userController := controllers.NewUserController(userRepo)
	photoController := controllers.NewPhotoController(photoRepo)
	commentController := controllers.NewCommentController(commentRepo)
	socialMediaController := controllers.NewSocialMediaController(socialMediaRepo)

	// Setup middleware
	authMiddleware := middlewares.NewAuthMiddleware(userRepo)

	// Setup routes
	api := router.Group("/api")
	{
		// Authentication routes
		api.POST("/register", authController.Register)
		api.POST("/login", authController.Login)

		// User routes
		api.GET("/users", userController.GetAll)
		api.GET("/users/:id", userController.GetOne)
		api.PUT("/users/:id", userController.Update)
		api.DELETE("/users/:id", userController.Delete)

		// Photo routes
		api.GET("/photos", photoController.GetAll)
		api.GET("/photos/:id", photoController.GetOne)
		api.POST("/photos", authMiddleware.Authenticate(), photoController.Create)
		api.PUT("/photos/:id", authMiddleware.Authenticate(), photoController.Update)
		api.DELETE("/photos/:id", authMiddleware.Authenticate(), photoController.Delete)

		// Comment routes
		api.GET("/comments", commentController.GetAll)
		api.GET("/comments/:id", commentController.GetOne)
		api.POST("/comments", authMiddleware.Authenticate(), commentController.Create)
		api.PUT("/comments/:id", authMiddleware.Authenticate(), commentController.Update)
		api.DELETE("/comments/:id", authMiddleware.Authenticate(), commentController.Delete)

		// Social media routes
		api.GET("/social-media", socialMediaController.GetAll)
		api.GET("/social-media/:id", socialMediaController.GetOne)
		api.POST("/social-media", authMiddleware.Authenticate(), socialMediaController.Create)
		api.PUT("/social-media/:id", authMiddleware.Authenticate(), socialMediaController.Update)
		api.DELETE("/social-media/:id", authMiddleware.Authenticate(), socialMediaController.Delete)
	}

	// Start server
	router.Run(":8080")
}
