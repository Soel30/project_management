package controllers

import (
	"net/http"
	"pm/domain"
	model "pm/models"
	"pm/repository"
	"pm/utils"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthController struct {
	DB *gorm.DB
}

func Login(ctx *gin.Context, DB *gorm.DB) {
	var user domain.User
	var token string
	var err error
	var server_config model.ServerConfig

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})

		return
	}

	result := repository.FindUserByUsername(DB, user, user.Username)
	if result.Error != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "Email or Password is wrong",
		})

		return
	}

	if !utils.CheckPasswordHash(user.Password, user.Password) {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "Email or Password is wrong",
		})

		return
	}

	claims := &model.Claims{
		User: user,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	jwt_secret := server_config.JWTSecret
	token, err = utils.GenerateToken(claims, jwt_secret)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal Server Error",
		})

		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"token": token,
		"user":  user,
	})
}

func Register(ctx *gin.Context, DB *gorm.DB) {
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
	result := DB.Create(&user)
	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal Server Error",
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "User Created",
		"user":    user,
	})

}

func RefreshToken(ctx *gin.Context, DB *gorm.DB) {
	var user domain.User
	var token string
	var err error
	var server_config model.ServerConfig

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})

		return
	}

	result := repository.FindUserByUsername(DB, user, user.Username)
	if result.Error != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "Email or Password is wrong",
		})

		return
	}

	if !utils.CheckPasswordHash(user.Password, user.Password) {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "Email or Password is wrong",
		})

		return
	}

	claims := &model.Claims{
		User: user,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	jwt_secret := server_config.JWTSecret
	token, err = utils.GenerateToken(claims, jwt_secret)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal Server Error",
		})

		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"token": token,
		"user":  user,
	})
}
