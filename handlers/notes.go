package handlers

import (
	"net/http"
	"strconv"

	"github.com/dev-tams/note-go/db"
	"github.com/dev-tams/note-go/models"
	"github.com/gin-gonic/gin"
)

func GetUserNotes(c *gin.Context) {


	var user models.User
	var notes models.Note

	//find user by id
	idParam := c.Param("id")
	id, _ := strconv.Atoi(idParam)

	if err := db.DB.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	//find user notes by user id
	if err := db.DB.Find(&notes).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch notes"})
		return
	}

	
	//return user notes

	c.JSON(http.StatusOK, gin.H{
		"message": "Notes from user",
		"data":    notes,
	})

}

func CreateUserNote(c *gin.Context) {
	//find user by id

	//create note for user
	//return created note
}

func UpdateUserNote(c *gin.Context) {
	//find user by id

	//find note by note id
	//update note for user
	//return updated note
}

func DeleteUserNote(c *gin.Context) {
	//find user by id

	//find note by note id
	//delete note for user
	//return deleted note
}
func GetNotes(c *gin.Context) {
	//find all notes
	//return all notes
}
