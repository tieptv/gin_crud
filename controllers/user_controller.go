package controllers

import (
	"gin_crud/models"
	"gin_crud/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type UserController struct {
	service services.UserService
}

func NewUserController(service services.UserService) *UserController {
	return &UserController{service: service}
}

// CreateUser godoc
func (uc *UserController) CreateUser(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := uc.service.Create(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Cannot create user"})
		return
	}

	c.JSON(http.StatusCreated, user)
}

// GetUsers godoc
func (uc *UserController) GetUsers(c *gin.Context) {
	users, err := uc.service.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Cannot fetch users"})
		return
	}

	c.JSON(http.StatusOK, users)
}

// GetUserByID godoc
func (uc *UserController) GetUserByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	user, err := uc.service.GetByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}

// UpdateUser godoc
func (uc *UserController) UpdateUser(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var input models.User
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := uc.service.Update(uint(id), &input)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}

// DeleteUser godoc
func (uc *UserController) DeleteUser(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := uc.service.Delete(uint(id)); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Deleted"})
}

func (uc *UserController) RegisterUserRoutes(r *gin.Engine) {
	users := r.Group("/users")
	{
		users.POST("", uc.CreateUser)
		users.GET("", uc.GetUsers)
		users.GET("/:id", uc.GetUserByID)
		users.PUT("/:id", uc.UpdateUser)
		users.DELETE("/:id", uc.DeleteUser)
	}
}
