package v1

import (
	"pm/app/controllers"
	"pm/app/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RoleRoutes(db *gorm.DB, router *gin.RouterGroup) {
	role := controllers.NewRoleController(db)

	middlew_chkt := middleware.CheckJwtAuth(db)
	router.GET("/roles", middlew_chkt, role.FindAll)
	router.GET("/roles/:id", middlew_chkt, role.FindById)
	router.POST("/roles", middlew_chkt, role.Create)
	router.PUT("/roles/:id", middlew_chkt, role.Update)
	router.DELETE("/roles/:id", middlew_chkt, role.Delete)
}
