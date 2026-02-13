package main

import (
	"gin_crud/db"
	_ "gin_crud/docs"
	"gin_crud/models"
	"gin_crud/routes"
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
	db.DB.AutoMigrate(&models.User{})

	routes.UserRoute(r)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run(":8080")
}
