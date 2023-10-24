package v1

import (
	"pm/app/controllers"
	"pm/app/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UserRoutes(db *gorm.DB, router *gin.RouterGroup) {
	user := controllers.NewUserController(db)

	middlew_chkt := middleware.CheckJwtAuth(db)
	router.GET("/users", middlew_chkt, user.FindAll)
	router.GET("/users/:id", middlew_chkt, user.FindById)
	router.POST("/users", middlew_chkt, user.Create)
	router.PUT("/users/:id", middlew_chkt, user.Update)
	router.DELETE("/users/:id", middlew_chkt, user.Delete)

}
