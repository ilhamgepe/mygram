package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ilhamgepe/mygram/cmd/api/handler"
	"github.com/ilhamgepe/mygram/cmd/api/middleware"
	"github.com/ilhamgepe/mygram/internal/repositories"
	"github.com/ilhamgepe/mygram/internal/services"
	"gorm.io/gorm"
)

func setupRoutes(engine *gin.Engine, db *gorm.DB) {

	// init repositories
	userRepository := repositories.NewUserRepository(db)
	photoRepository := repositories.NewPhotoRepository(db, userRepository)
	commentRepository := repositories.NewCommentRepository(db)
	socialmediaRepository := repositories.NewSocialMediaRepository(db)

	// init services
	userService := services.NewUserService(userRepository)
	photoService := services.NewPhotoService(photoRepository)
	commentService := services.NewCommentService(commentRepository)
	socialmediaService := services.NewSocialMediaService(socialmediaRepository)

	// init handler
	userHandler := handler.NewUserHandler(userService)
	photoHandler := handler.NewPhotoHandler(photoService)
	commentHandler := handler.NewCommentHandler(commentService)
	socialMediaHandler := handler.NewSocialMediaHandler(socialmediaService)

	user := engine.Group("/users")
	{
		user.POST("/register", userHandler.Register)
		user.POST("/login", userHandler.Login)

		user.Use(middleware.WithAuth())
		user.PUT("/:id", userHandler.Update)
		user.DELETE("/", userHandler.Delete)
	}

	photo := engine.Group("/photos")
	{
		photo.Use(middleware.WithAuth())
		photo.POST("/", photoHandler.AddPhoto)
		photo.GET("/", photoHandler.GetPhotos)
		photo.PUT("/:id", photoHandler.UpdatePhoto)
		photo.DELETE("/:id", photoHandler.DeletePhoto)
	}

	comment := engine.Group("/comments")
	{
		comment.Use(middleware.WithAuth())
		comment.POST("/", commentHandler.AddComment)
		comment.GET("/", commentHandler.GetComments)
		comment.PUT("/:id", commentHandler.UpdateComment)
		comment.DELETE("/:id", commentHandler.DeleteComment)
	}

	socialMedia := engine.Group("/socialmedias")
	{
		socialMedia.Use(middleware.WithAuth())
		socialMedia.POST("/", socialMediaHandler.AddSocialMedia)
		socialMedia.GET("/", socialMediaHandler.GetSocialMedias)
		socialMedia.PUT("/:id", socialMediaHandler.UpdateSocialMedia)
		socialMedia.DELETE("/:id", socialMediaHandler.DeleteSocialMedia)
	}

	engine.GET("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "Hello World2")
	})
}
