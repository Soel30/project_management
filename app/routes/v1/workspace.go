package v1

import (
	"pm/app/controllers"
	"pm/app/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func WorkspaceRoutes(db *gorm.DB, router *gin.RouterGroup) {
	workspace := controllers.NewWorkspaceController(db)

	middlew_chkt := middleware.CheckJwtAuth(db)
	router.GET("/workspaces", middlew_chkt, workspace.FindAll)
	router.GET("/workspaces/:id", middlew_chkt, workspace.FindById)
	router.POST("/workspaces", middlew_chkt, workspace.Create)
	router.PUT("/workspaces/:id", middlew_chkt, workspace.Update)
	router.DELETE("/workspaces/:id", middlew_chkt, workspace.Delete)
	router.POST("/workspaces/add_user/:workspace_id", middlew_chkt, workspace.AddUserToWorkspace)
	router.DELETE("/workspaces/remove_user/:workspace_id/:user_id", middlew_chkt, workspace.RemoveUserFromWorkspace)
}
