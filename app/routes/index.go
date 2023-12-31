package routes

import (
	v1 "pm/app/routes/v1"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Routes(db *gorm.DB, engine *gin.Engine) {
	api := engine.Group("/api/v1")
	v1.UserRoutes(db, api)
	v1.RoleRoutes(db, api)
	v1.WorkspaceRoutes(db, api)
	v1.CategoryRoutes(db, api)
}
