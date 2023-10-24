package controllers

import (
	"net/http"
	"pm/domain"
	"pm/repository"
	"pm/utils"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserController struct {
	DB *gorm.DB
}

func NewUserController(db *gorm.DB) UserController {
	return UserController{
		DB: db,
	}
}

func (c *UserController) FindAll(ctx *gin.Context) {
	var users domain.User
	pagination := utils.GenerateUserPagination(ctx)

	userLists, err := repository.GetAllUser(&users, pagination, c.DB, "")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal Server Error",
		})
		return
	}

	ctx.JSON(http.StatusOK, userLists)
}

func (c *UserController) FindById(ctx *gin.Context) {
	var user domain.User
	result := c.DB.First(&user, ctx.Param("id")).Preload("Workspaces")
	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal Server Error",
		})
		return
	}

	ctx.JSON(http.StatusOK, user)
}

func (c *UserController) Create(ctx *gin.Context) {
	var user domain.User

	// handle photo upload from form-data
	file, _ := ctx.FormFile("photo")
	if file != nil {
		photo, err := utils.UploadFile(file)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": "Internal Server Error",
			})
			return
		}
		user.Photo = photo
	}

	// handle form-data
	err := ctx.ShouldBind(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad Request",
		})
		return
	}

	// create user
	// has password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal Server Error",
		})
		return
	}

	user.Password = string(hashedPassword)
	result := c.DB.Create(&user)
	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal Server Error",
		})
		return
	}

	ctx.JSON(http.StatusCreated, user)
}

func (c *UserController) Update(ctx *gin.Context) {
	var user domain.User
	user_id := ctx.Param("id")
	var hashedPassword []byte

	// find user by id
	result := c.DB.First(&user, user_id)
	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal Server Error",
		})
		return
	}

	// handle photo upload from form-data
	file, _ := ctx.FormFile("photo")
	if file != nil {
		photo, err := utils.UploadFile(file)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": "Internal Server Error",
			})
			return
		}
		user.Photo = photo
	}

	err := ctx.ShouldBind(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad Request",
		})
		return
	}

	if user.Password != "" {
		hashedPassword, err = bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": "Internal Server Error",
			})

			return
		}
	} else {
		hashedPassword = []byte(user.Password)
	}

	user.Password = string(hashedPassword)
	result = result.Updates(&user)
	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal Server Error",
		})
		return
	}

	ctx.JSON(http.StatusOK, user)
}

func (c *UserController) Delete(ctx *gin.Context) {
	var user domain.User
	user_id := ctx.Param("id")

	// find user by id
	result := c.DB.First(&user, user_id)
	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal Server Error",
		})
		return
	}

	// delete user
	result = c.DB.Delete(&user, user_id)
	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal Server Error",
		})
		return
	}
}
