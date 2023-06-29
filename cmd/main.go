package main

import (
	"log"
	"net/http"
	"sharing_vision/app/routes"
	"sharing_vision/config/db"
	"sharing_vision/config/env"
	"sharing_vision/domain"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	env.NewEnv(".env")

	routes.Routes(db.NewDB(env.Config).DB, router)
	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"code": "PAGE_NOT_FOUND", "message": "404 page not found"})
	})

	router.NoMethod(func(c *gin.Context) {
		c.JSON(http.StatusMethodNotAllowed, gin.H{"code": "METHOD_NOT_ALLOWED", "message": "405 method not allowed"})
	})

	dbBase := db.NewDB(env.Config).DB
	// Migrate the schema
	err := dbBase.Debug().Migrator().AutoMigrate(&domain.Post{})
	if err != nil {
		panic(err)
	}

	if err := router.Run(env.Config.Host + ":" + env.Config.Port); err != nil {
		log.Fatal("main : error starting server", err)
	}
}
