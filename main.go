package main

import (
	"log"

	"library/config"
	"library/controllers"
	"library/db"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.Load()
	database := db.Open(cfg.DatabaseURL)
	defer database.Close()

	auth := controllers.AuthController{DB: database}
	cats := controllers.CategoriesController{DB: database}
	books := controllers.BooksController{DB: database}

	r := gin.Default()

	r.POST("/api/users/login", auth.Login)

	api := r.Group("/api")
	api.Use(controllers.JWTAuth())

	api.GET("/categories", cats.List)
	api.POST("/categories", cats.Create)
	api.GET("/categories/:id", cats.Detail)
	api.DELETE("/categories/:id", cats.Delete)
	api.GET("/categories/:id/books", cats.BooksByCategory)

	api.GET("/books", books.List)
	api.POST("/books", books.Create)
	api.DELETE("/books/:id", books.Delete)
	api.GET("/books/:id", books.Detail)

	log.Printf("Server running on port %s", cfg.Port)
	if err := r.Run(":" + cfg.Port); err != nil {
		log.Fatal(err)
	}
}
