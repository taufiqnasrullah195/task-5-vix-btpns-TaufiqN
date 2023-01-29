package router

import (
	"github.com/gin-gonic/gin"
	"github.com/taufiqnasrullah195/task-5-vix-btpns-TaufiqNashrullah/tree/main/controllers"
	"github.com/taufiqnasrullah195/task-5-vix-btpns-TaufiqNashrullah/tree/main/middlewares"
)

func StartApp() *gin.Engine {
	r := gin.Default()

	userRouter := r.Group("/users")
	{
		userRouter.POST("/register", controllers.UserRegister)
		userRouter.POST("/login", controllers.UserLogin)
		userRouter.PUT("/:userId", controllers.UserUpdate)
		userRouter.DELETE("/:userId", controllers.UserDelete)
	}

	photoRouter := r.Group("/photos")
	{
		photoRouter.Use(middlewares.Authentication())

		photoRouter.POST("/", controllers.CreatePhoto)
		photoRouter.GET("/", controllers.ListPhoto)
		photoRouter.PUT("/:photoId", middlewares.PhotoAuthorization(), controllers.UpdatePhoto)
		photoRouter.DELETE("/:photoId", middlewares.PhotoAuthorization(), controllers.DeletePhoto)
	}

	return r
}
