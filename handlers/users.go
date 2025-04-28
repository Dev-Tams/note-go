package handlers

import (
	"fmt"
	"net/http"
	"slices"
	"strconv"

	"github.com/dev-tams/note-go/models"
	"github.com/gin-gonic/gin"
)

var users []models.User

func GetUsers(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"message": "All Users",
		"data":    users,
	})
}

func GetUserByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid user ID",
		})
		return
	}
	for _, user := range users {
		if int(user.ID) == id {
			c.JSON(http.StatusOK, gin.H{
				"data": user,
			})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{
		"error": "User not found",
	})
}

func CreateUser(c *gin.Context) {

	var user models.User

	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid JSON format",
		})
		return
	}
	user.ID = uint(len(users) + 1)
	users = append(users, user)

	fmt.Printf("User: %+v\n", user)

	// Respond to the client
	c.JSON(http.StatusCreated, user)
}

func DeleteUser(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
        return
    }

	for index, user := range users {
		if int(user.ID) == id {

			users = slices.Delete(users, index, index+1)

			c.Status(http.StatusNoContent)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{
		"error": "User not found",
	})
}
