package routes

import (
	v1 "sharing_vision/app/routes/v1"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Routes(db *gorm.DB, engine *gin.Engine) {
	api := engine.Group("/api/v1")
	v1.ArticleRoutes(db, api)
}
