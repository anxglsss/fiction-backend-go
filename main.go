package main

import (
	"log"

	"fiction-turnament/config"
	"fiction-turnament/db"
	"fiction-turnament/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.Load()

	if err := db.Init(cfg.DatabaseURL); err != nil {
		log.Fatalf("Ошибка подключения к БД: %v", err)
	}

	r := gin.Default()
	r.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", cfg.CORSOrigin)
		c.Header("Access-Control-Allow-Methods", "GET, POST")
		c.Header("Access-Control-Allow-Headers", "Content-Type")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	api := r.Group("/api")
	{
		api.POST("/users", handlers.CreateUser)
		api.GET("/users", handlers.ListUsers)
		api.GET("/fictions", handlers.GetFictions)
		api.POST("/tournaments", handlers.CreateTournament)
		api.GET("/tournaments", handlers.ListTournaments)
		api.GET("/tournaments/:id", handlers.GetTournament)
		api.POST("/tournaments/:id/matches/:matchId/vote", handlers.VoteMatch)
	}

	log.Printf("Сервер запущен на http://localhost:%s", cfg.ServerPort)
	log.Printf("База данных:%s", cfg.DatabaseURL)
	if err := r.Run(":" + cfg.ServerPort); err != nil {
		log.Fatalf("Ошибка сервера: %v", err)
	}
}
