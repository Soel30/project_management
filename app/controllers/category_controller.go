package controllers

import (
	"net/http"
	"pm/domain"
	"pm/repository"
	"pm/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CategoryController struct {
	DB *gorm.DB
}

func NewCategoryController(db *gorm.DB) CategoryController {
	return CategoryController{
		DB: db,
	}
}

func (c *CategoryController) FindAll(ctx *gin.Context) {
	var categories domain.Category
	pagination := utils.GenerateCategoryPagination(ctx)

	categoryLists, err := repository.GetAllCategory(&categories, pagination, c.DB, "")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal Server Error",
		})
		return
	}

	ctx.JSON(http.StatusOK, categoryLists)
}

func (c *CategoryController) FindById(ctx *gin.Context) {
	var category domain.Category
	id := ctx.Param("id")

	categoryLists := c.DB.First(&category, id)
	if categoryLists.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal Server Error",
		})
		return
	}

	ctx.JSON(http.StatusOK, category)

}

func (c *CategoryController) Create(ctx *gin.Context) {
	var category domain.Category

	if err := ctx.ShouldBindJSON(&category); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad Request",
		})
		return
	}

	categoryLists := c.DB.Create(&category)
	if categoryLists.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal Server Error",
		})
		return
	}
	ctx.JSON(http.StatusOK, category)
}

func (c *CategoryController) Update(ctx *gin.Context) {
	var category domain.Category
	id := ctx.Param("id")

	categoryLists := c.DB.First(&category, id)
	if categoryLists.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal Server Error",
		})
		return
	}

	if err := ctx.ShouldBindJSON(&category); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad Request",
		})
		return
	}

	categoryLists = c.DB.Save(&category)
	if categoryLists.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal Server Error",
		})
		return
	}
	ctx.JSON(http.StatusOK, category)
}

func (c *CategoryController) Delete(ctx *gin.Context) {
	var category domain.Category
	id := ctx.Param("id")

	categoryLists := c.DB.First(&category, id)
	if categoryLists.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal Server Error",
		})
		return
	}

	categoryLists = c.DB.Delete(&category, id)
	if categoryLists.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal Server Error",
		})
		return
	}
	ctx.JSON(http.StatusOK, category)
}
