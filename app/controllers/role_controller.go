package controllers

import (
	"net/http"
	"pm/domain"
	"pm/repository"
	"pm/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type RoleController struct {
	DB *gorm.DB
}

func NewRoleController(db *gorm.DB) RoleController {
	return RoleController{
		DB: db,
	}
}

func (c *RoleController) FindAll(ctx *gin.Context) {
	var roles domain.Role
	pagination := utils.GenerateRolePagination(ctx)

	roleLists, err := repository.GetAllRole(&roles, pagination, c.DB, "")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal Server Error",
		})
		return
	}

	ctx.JSON(http.StatusOK, roleLists)
}

func (c *RoleController) FindById(ctx *gin.Context) {
	var role domain.Role
	result := c.DB.First(&role, ctx.Param("id"))
	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal Server Error",
		})
		return
	}

	ctx.JSON(http.StatusOK, role)
}

func (c *RoleController) Create(ctx *gin.Context) {
	var role domain.Role

	// bind json
	if err := ctx.ShouldBindJSON(&role); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})

		return
	}

	result := c.DB.Create(&role)
	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal Server Error",
		})
		return
	}

	ctx.JSON(http.StatusCreated, role)
}

func (c *RoleController) Update(ctx *gin.Context) {
	var role domain.Role
	result := c.DB.First(&role, ctx.Param("id"))
	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal Server Error",
		})
		return
	}

	// bind json
	if err := ctx.ShouldBindJSON(&role); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})

		return
	}

	result = c.DB.Save(&role)
	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal Server Error",
		})
		return
	}

	ctx.JSON(http.StatusOK, role)
}

func (c *RoleController) Delete(ctx *gin.Context) {
	var role domain.Role
	role_id := ctx.Param("id")

	// find role by id
	result := c.DB.First(&role, role_id)
	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal Server Error",
		})
		return
	}

	// delete role
	result = c.DB.Delete(&role, role_id)
	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal Server Error",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Role deleted successfully",
	})
}
