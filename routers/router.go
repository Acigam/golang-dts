package routers

import (
	"final-project-acgm/controllers"
	"final-project-acgm/middlewares"

	"github.com/gin-gonic/gin"
	// _ "final-project-acgm/docs"
	// ginSwagger "github.com/swaggo/gin-swagger"
	// swaggerfiles "github.com/swaggo/files"
)

func StartApp() *gin.Engine {
	r := gin.Default()
	apiRouter := r.Group("/api")

	userRouter := apiRouter.Group("/users")
	{
		userRouter.POST("/register", controllers.UserRegister) // Register a new user
		userRouter.POST("/login", controllers.UserLogin)       // Login user
	}

	lockerRouter := apiRouter.Group("/lockers")
	{
		lockerRouter.Use(middlewares.Authentication())
		lockerRouter.POST("", controllers.LockerCreate)             // Create a new locker
		lockerRouter.GET("", controllers.LockerList)                // Get all lockers
		lockerRouter.PUT("/:lockerId", controllers.LockerUpdate)    // Update a locker
		lockerRouter.DELETE("/:lockerId", controllers.LockerDelete) // Delete a locker
	}

	lockerRentRouter := apiRouter.Group("/locker-rents")
	{
		lockerRentRouter.Use(middlewares.Authentication())
		lockerRentRouter.POST("", controllers.LockerRentCreate) // Create a new locker rent
		lockerRentRouter.GET("", controllers.LockerRentList)    // Get all locker rents
	}

	return r
}
