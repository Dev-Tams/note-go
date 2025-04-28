package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

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
	c.JSON(http.StatusFound, gin.H{
		"message": "Get User By ID",
	})
}

func CreateUser(c *gin.Context) {
	// Read the request body

	fileBytes, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request body",
		})
		return
	}
	// Create a User struct to store the unmarshaled data
	var user models.User
	// Unmarshal the JSON data into the User struct

	err = json.Unmarshal(fileBytes, &user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid JSON format",
		})
	}
	users = append(users, user)
	// Print the user data (or do something else with it)

	fmt.Printf("User: %+v\n", user)

	// Respond to the client
	c.JSON(http.StatusCreated, user)
}

func DeleteUser(c *gin.Context) {
	c.JSON(http.StatusMovedPermanently, gin.H{
		"message": "Delete User",
	})
}
