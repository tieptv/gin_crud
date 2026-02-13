package bootstrap

import (
	"github.com/gin-gonic/gin"
	"os/user"
)

func Run() {
	db := database.New()

	db.AutoMigrate(&user.User{})

	r := gin.Default()

	user.RegisterModule(r, db)

	r.Run(":8080")
}
