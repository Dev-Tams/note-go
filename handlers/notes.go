package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/dev-tams/note-go/db"
	"github.com/dev-tams/note-go/models"
	"github.com/gin-gonic/gin"
)

func GetUserNotes(c *gin.Context) {
	var user models.User
	var notes []models.Note

	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	if err := db.DB.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	if err := db.DB.Where("user_id = ?", id).Find(&notes).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch notes"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Notes from user",
		"data":    notes,
	})
}
func GetUserNoteById(c *gin.Context) {
	var user models.User
	var note models.Note

	// Get user ID and note ID from URL
	userIDParam := c.Param("id")
	noteIDParam := c.Param("noteId")

	userID, err := strconv.Atoi(userIDParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	noteID, err := strconv.Atoi(noteIDParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid note ID"})
		return
	}

	if err := db.DB.First(&user, userID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	if err := db.DB.Where("id = ? AND user_id = ?", noteID, userID).First(&note).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Note not found for this user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": note,
	})
}

func CreateUserNote(c *gin.Context) {
	var note models.Note

	if err := c.ShouldBindJSON(&note); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON format", "details": err.Error()})
		return
	}

	var user models.User
	if err := db.DB.First(&user, note.UserID).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found", "details": err.Error()})
		return
	}

	// Create the note
	if err := db.DB.Create(&note).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create note", "details": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, note)
}

func UpdateUserNote(c *gin.Context) {
	var note models.Note
	var user models.User

	// Get user ID and note ID from URL parameters
	userIdParam := c.Param("id")
	noteIdParam := c.Param("noteId")

	userID, err := strconv.Atoi(userIdParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid user id"})
		return
	}
	noteID, err := strconv.Atoi(noteIdParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid note id"})
		return
	}

	if err := db.DB.First(&user, userID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	if err := db.DB.Where("id = ? AND user_id = ?", noteID, userID).First(&note).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Note not found for this user"})
		return
	}

	// Update note for user
	var updatedNote models.Note
	if err := c.BindJSON(&updatedNote); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	note.Title = updatedNote.Title
	note.Content = updatedNote.Content

	if err := db.DB.Save(&note).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update note"})
		return
	}

	// Return updated note
	c.JSON(http.StatusOK, gin.H{"message": "Note updated", "data": note})
}

// func UpdateUserNote(c *gin.Context) {
//     c.JSON(200, gin.H{"message": "hit update user note"})
// }

func DeleteUserNote(c *gin.Context) {
	log.Println("Handler hit!")
	//find user by id

	//find note by note id
	//delete note for user
	//return deleted note
	log.Println("Returning JSON response")
	c.JSON(http.StatusOK, gin.H{"status": "hello world"})
}
func GetNotes(c *gin.Context) {
	//find all notes
	//return all notes
}
