package v1

import (
	"pm/app/controllers"
	"pm/app/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CategoryRoutes(db *gorm.DB, router *gin.RouterGroup) {
	category := controllers.NewCategoryController(db)

	middlew_chkt := middleware.CheckJwtAuth(db)
	router.GET("/categories", middlew_chkt, category.FindAll)
	router.GET("/categories/:id", middlew_chkt, category.FindById)
	router.POST("/categories", middlew_chkt, category.Create)
	router.PUT("/categories/:id", middlew_chkt, category.Update)
	router.DELETE("/categories/:id", middlew_chkt, category.Delete)

}
