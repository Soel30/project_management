package controllers

import (
	"net/http"
	"sharing_vision/domain"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type PostController struct {
	DB *gorm.DB
}

func NewPostController(db *gorm.DB) PostController {
	return PostController{
		DB: db,
	}
}

func (c *PostController) FindAll(ctx *gin.Context) {
	var posts []domain.Post
	limit_query := ctx.Query("limit")
	if limit_query == "" {
		limit_query = "-1"
	}
	offset_query := ctx.Query("offset")
	if offset_query == "" {
		offset_query = "-1"
	}

	limitInt, _ := strconv.Atoi(limit_query)
	offsetInt, _ := strconv.Atoi(offset_query)
	result := c.DB.Limit(limitInt).Offset(offsetInt).Find(&posts)
	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal Server Error",
		})
		return
	}
	ctx.JSON(http.StatusOK, posts)
}

func (c *PostController) FindById(ctx *gin.Context) {
	var post domain.Post
	result := c.DB.First(&post, ctx.Param("id"))
	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal Server Error",
		})
		return
	}
	ctx.JSON(http.StatusOK, post)
}

func (c *PostController) Create(ctx *gin.Context) {
	var post domain.Post

	if err := ctx.ShouldBindJSON(&post); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := c.DB.Create(&post)
	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal Server Error",
		})
		return
	}
	ctx.JSON(http.StatusCreated, post)
}

func (c *PostController) Update(ctx *gin.Context) {
	var post domain.Post
	result := c.DB.First(&post, ctx.Param("id"))
	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal Server Error",
		})
		return
	}
	if err := ctx.ShouldBindJSON(&post); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result = c.DB.Save(&post)
	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal Server Error",
		})
		return
	}
	ctx.JSON(http.StatusOK, post)
}

func (c *PostController) Delete(ctx *gin.Context) {
	result := c.DB.Delete(&domain.Post{}, ctx.Param("id"))
	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal Server Error",
		})
		return
	}
	ctx.JSON(http.StatusNoContent, nil)
}

func (c *PostController) FindByLimitOffset(ctx *gin.Context) {
	var posts []domain.Post
	limit := ctx.Param("limit")
	offset := ctx.Param("offset")
	limitInt, _ := strconv.Atoi(limit)
	offsetInt, _ := strconv.Atoi(offset)
	result := c.DB.Limit(limitInt).Offset(offsetInt).Find(&posts)
	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal Server Error",
		})
		return
	}
	ctx.JSON(http.StatusOK, posts)
}
