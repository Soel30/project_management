package v1

import (
	"sharing_vision/app/controllers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ArticleRoutes(db *gorm.DB, router *gin.RouterGroup) {
	article := controllers.NewPostController(db)

	router.GET("/articles", article.FindAll)
	router.GET("/articles/:id", article.FindById)
	router.POST("/articles", article.Create)
	router.PUT("/articles/:id", article.Update)
	router.DELETE("/articles/:id", article.Delete)

}
