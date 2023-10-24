package controllers

import (
	"net/http"
	"pm/domain"
	"pm/repository"
	"pm/utils"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type WorkspaceController struct {
	DB *gorm.DB
}

type WorkspaceRequest struct {
	Name   string `json:"name" binding:"required"`
	Color  string `json:"color" binding:"required"`
	UserId uint   `json:"user_id" binding:"required"`
	RoleId int    `json:"role_id" binding:"required"`
}

type AddUserToWorkspaceRequest struct {
	UserId uint `json:"user_id" binding:"required"`
	RoleId int  `json:"role_id" binding:"required"`
}

func NewWorkspaceController(db *gorm.DB) WorkspaceController {
	return WorkspaceController{
		DB: db,
	}
}

func (c *WorkspaceController) FindAll(ctx *gin.Context) {
	var workspaces domain.Workspace
	pagination := utils.GenerateWorkspacePagination(ctx)

	workspaceLists, err := repository.GetAllWorkspace(&workspaces, pagination, c.DB, "")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal Server Error",
		})
		return
	}

	ctx.JSON(http.StatusOK, workspaceLists)
}

func (c *WorkspaceController) FindById(ctx *gin.Context) {
	var workspace domain.Workspace
	result := c.DB.First(&workspace, ctx.Param("id")).Preload("Users")
	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal Server Error",
		})
		return
	}

	ctx.JSON(http.StatusOK, workspace)
}

func (c *WorkspaceController) Create(ctx *gin.Context) {
	var workspace WorkspaceRequest
	var user domain.User
	var userWorkspace domain.UserWorkspace

	if err := ctx.ShouldBindJSON(&workspace); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	result := c.DB.Where("id = ?", workspace.UserId).First(&user)
	if result.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "User not found",
		})
		return
	}

	workspaceData := domain.Workspace{
		Name:  workspace.Name,
		Color: workspace.Color,
	}

	workspaceResult := c.DB.Create(&workspaceData)

	userWorkspace.UserId = workspace.UserId
	userWorkspace.WorkspaceId = int(workspaceResult.RowsAffected)
	userWorkspace.RoleId = workspace.RoleId

	userWorkspaceResult := c.DB.Create(&userWorkspace)
	if userWorkspaceResult.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal Server Error",
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Workspace created successfully",
	})

}

func (c *WorkspaceController) Update(ctx *gin.Context) {
	var workspace domain.Workspace
	var workspaceRequest WorkspaceRequest

	if err := ctx.ShouldBindJSON(&workspaceRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	result := c.DB.First(&workspace, ctx.Param("id"))
	if result.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Workspace not found",
		})
		return
	}

	workspace.Name = workspaceRequest.Name
	workspace.Color = workspaceRequest.Color

	c.DB.Save(&workspace)

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Workspace updated successfully",
	})
}

func (c *WorkspaceController) Delete(ctx *gin.Context) {
	var workspace domain.Workspace

	result := c.DB.First(&workspace, ctx.Param("id"))
	if result.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Workspace not found",
		})
		return
	}

	c.DB.Delete(&workspace)

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Workspace deleted successfully",
	})
}

func (c *WorkspaceController) AddUserToWorkspace(ctx *gin.Context) {
	var user domain.User
	var userWorkspace domain.UserWorkspace
	var addUserToWorkspaceRequest AddUserToWorkspaceRequest

	workspace_id, _ := strconv.Atoi(ctx.Param("workspace_id"))

	// bind json request to struct
	if err := ctx.ShouldBindJSON(&addUserToWorkspaceRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	// check if user exist
	result := c.DB.Where("id = ?", addUserToWorkspaceRequest.UserId).First(&user)
	if result.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "User not found",
		})
		return
	}

	userWorkspace.UserId = user.ID
	userWorkspace.WorkspaceId = workspace_id
	userWorkspace.RoleId = addUserToWorkspaceRequest.RoleId

	userWorkspaceResult := c.DB.Create(&userWorkspace)
	if userWorkspaceResult.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal Server Error",
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "User added to workspace successfully",
	})
}

func (c *WorkspaceController) RemoveUserFromWorkspace(ctx *gin.Context) {
	var userWorkspace domain.UserWorkspace

	workspace_id, _ := strconv.Atoi(ctx.Param("workspace_id"))

	result := c.DB.Where("workspace_id = ? AND user_id = ?", workspace_id, ctx.Param("user_id")).First(&userWorkspace)
	if result.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "User not found in workspace",
		})
		return
	}

	c.DB.Delete(&userWorkspace)

	ctx.JSON(http.StatusOK, gin.H{
		"message": "User removed from workspace successfully",
	})
}
