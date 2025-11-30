package routes

import (
	"ibnufth/backend-api/controller"
	"ibnufth/backend-api/middlewares"
	"ibnufth/backend-api/version"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	// initialize a new Gin router
	router := gin.Default()

	// set up CORS
	router.Use(cors.New(cors.Config{
		AllowOrigins:  []string{"*"},
		AllowMethods:  []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:  []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders: []string{"Content-Length"},
	}))

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Hello World",
			"version": version.GetVersion(),
		})
	})

	// route register
	router.POST("/api/register", controller.Register)

	// route login
	router.POST("/api/login", controller.Login)

	// route find user
	router.GET("/api/users", middlewares.AuthMiddleware(), controller.FindUser)

	// route create user
	router.POST("/api/users", middlewares.AuthMiddleware(), controller.CreateUser)

	// route find user by id
	router.GET("/api/users/:id", middlewares.AuthMiddleware(), controller.FindUserById)

	// route update user
	router.PUT("/api/users/:id", middlewares.AuthMiddleware(), controller.UpdateUser)

	// route delete user
	router.DELETE("/api/users/:id", middlewares.AuthMiddleware(), controller.DeleteUser)

	return router
}
