package controllers

import (
	"net/http"
	"sharing_vision/domain"
	"sharing_vision/repository"
	"sharing_vision/utils"
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
	var posts domain.Post
	pagination := utils.GeneratePaginationFromRequest(ctx)
	status_query := ctx.Query("status")

	if status_query == "" {
		postLists, err := repository.GetAllPost(&posts, pagination, c.DB, "")
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": "Internal Server Error",
			})
			return
		}

		ctx.JSON(http.StatusOK, postLists)
	} else {
		postLists, err := repository.GetAllPost(&posts, pagination, c.DB, status_query)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": "Internal Server Error",
			})
			return
		}

		ctx.JSON(http.StatusOK, postLists)
	}

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
	var result *gorm.DB
	var errorArray []string

	if err := ctx.ShouldBindJSON(&post); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if post.Status == "Draft" {
		result = c.DB.Save(&post)
	} else if post.Status == "Publish" {
		if post.Title == "" {
			errorArray = append(errorArray, "Title cannot be empty")
		}
		if post.Content == "" {
			errorArray = append(errorArray, "Content cannot be empty")
		}
		if post.Category == "" {
			errorArray = append(errorArray, "Category cannot be empty")
		}
		if len(post.Title) < 20 {
			errorArray = append(errorArray, "Title minimum length is 20")
		}
		if len(post.Content) < 200 {
			errorArray = append(errorArray, "Content minimum length is 200")
		}
		if len(post.Category) < 3 {
			errorArray = append(errorArray, "Category minimum length is 3")
		}

		if len(errorArray) > 0 {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": errorArray,
			})
			return
		}
		result = c.DB.Save(&post)

	} else {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Status must be Draft or Publish",
		})
		return
	}
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
	// result := c.DB.Delete(&domain.Post{}, ctx.Param("id"))
	// change status to Thrash
	var post domain.Post
	result := c.DB.First(&post, ctx.Param("id"))
	post.Status = "Thrash"
	result = c.DB.Save(&post)
	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal Server Error",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Post deleted successfully",
	})
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
