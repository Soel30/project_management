package main

import (
	"log"
	"net/http"
	"pm/app/routes"
	"pm/config/db"
	"pm/config/env"
	"pm/domain"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	// set
	env.NewEnv(".env")
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"PUT", "PATCH", "POST", "GET", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization", "accept", "origin", "Cache-Control", "X-Requested-With", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Access-Control-Allow-Methods", "Access-Control-Allow-Credentials"},
		AllowCredentials: true,
	}))

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
