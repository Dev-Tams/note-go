package main

import (
	"log"
	"github.com/dev-tams/note-go/db"
	"github.com/dev-tams/note-go/handlers"

	// "net/http"
	"github.com/gin-gonic/gin"
)

func main() {
	
	db.Init()
	r := gin.Default()
	RegisterRoutes(r)

	
	log.Print("listening on Port 8000")
	errr := r.Run(":8000")
	if errr != nil {
		log.Fatal("Server Run Failed:", errr)
	}

}

func RegisterRoutes(r *gin.Engine) {
	r.GET("/hello", handlers.Hello)

	r.GET("/users", handlers.GetUsers)
	r.GET("/users/:id", handlers.GetUserByID)
	r.POST("/users", handlers.CreateUser)
	r.PUT("/users/:id", handlers.UpdateUser)
	r.DELETE("/users/:id", handlers.DeleteUser)

	r.GET("/users/:id/notes", handlers.GetUserNotes)
	r.POST("/users/:id/notes", handlers.CreateUserNote)
	r.PUT("/users/:id/notes/:noteId", handlers.UpdateUserNote)
	r.DELETE("/users/:id/notes/:noteId", handlers.DeleteUserNote)
	r.GET("/notes", handlers.GetNotes)
}
