package main

import (
	"gin_crud/controllers"
	"gin_crud/db"
	_ "gin_crud/docs"
	"gin_crud/models"
	"gin_crud/repositories"
	"gin_crud/services"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	_ "github.com/swaggo/gin-swagger"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Simple CRUD API
// @version 1.0
// @description Basic Gin CRUD with Swagger
// @host localhost:8080
// @BasePath /
func main() {

	r := gin.Default()

	db.ConnectDB()

	// Auto migrate
	err := db.DB.AutoMigrate(&models.User{})
	if err != nil {
		return
	}

	userRepo := repositories.NewUserRepository(db.DB)
	userService := services.NewUserService(userRepo)
	userController := controllers.NewUserController(userService)

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: false,
	}))
	userController.RegisterUserRoutes(r)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run(":8080")
}
