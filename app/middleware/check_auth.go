package middleware

import (
	"net/http"
	model "pm/models"
	"pm/repository"
	"pm/utils"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

func CheckJwtAuth(DB *gorm.DB) gin.HandlerFunc {
	var server_config model.ServerConfig
	return func(ctx *gin.Context) {
		var token string
		var err error
		var claims model.Claims

		header := ctx.GetHeader("Authorization")
		if header == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"message": "Unauthorized",
			})

			ctx.Abort()

			return
		}

		splitToken := strings.Split(header, "Bearer ")
		if len(splitToken) != 2 {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"message": "Unauthorized",
			})

			ctx.Abort()

			return
		}

		jwt_secret := server_config.JWTSecret
		token = splitToken[1]
		_, err = jwt.ParseWithClaims(token, &claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(jwt_secret), nil
		})
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"message": "Unauthorized",
			})

			ctx.Abort()

			return
		}

		result := repository.FindUserByUsername(DB, claims.User, claims.User.Username)
		if result.Error != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"message": "Unauthorized",
			})

			ctx.Abort()

			return
		}

		if !utils.CheckPasswordHash(claims.User.Password, claims.User.Password) {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"message": "Unauthorized",
			})

			ctx.Abort()

			return
		}

		ctx.Next()
	}
}
