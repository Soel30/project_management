package v1

import (
	"pm/app/controllers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UserRoutes(db *gorm.DB, router *gin.RouterGroup) {
	user := controllers.NewUserController(db)

	router.GET("/users", user.FindAll)
	router.GET("/users/:id", user.FindById)
	router.POST("/users", user.Create)
	router.PUT("/users/:id", user.Update)
	router.DELETE("/users/:id", user.Delete)

}
