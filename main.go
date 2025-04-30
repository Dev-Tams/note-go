package main

import (
	"log"

	"github.com/dev-tams/note-go/db"
	"github.com/dev-tams/note-go/handlers"
	"github.com/joho/godotenv"

	// "net/http"
	"github.com/gin-gonic/gin"
)

func main() {
	db.Init()
	r := gin.Default()

	godotenv.Load()

	RegisterRoutes(r)

	log.Print("listening on Port 8000")
	err := r.Run(":8000")
	if err != nil {
		log.Fatal("Server Run Failed:", err)
	}

}

func RegisterRoutes(r *gin.Engine) {
	r.GET("/hello", handlers.Hello)

	r.GET("/users", handlers.GetUsers)
	r.GET("/users/:id", handlers.GetUserByID)
	r.POST("/users", handlers.CreateUser)
	r.DELETE("/users/:id", handlers.DeleteUser)
}
