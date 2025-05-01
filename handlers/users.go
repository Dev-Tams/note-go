package handlers

import (
	"net/http"
	"strconv"
	"github.com/dev-tams/note-go/db"
	"github.com/dev-tams/note-go/models"
	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {

	var users []models.User
	if err := db.DB.Find(&users).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch users"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "All Users",
		"data":    users,
	})
}

func GetUserByID(c *gin.Context) {
	var user models.User

	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid user ID",
		})
		return
	}

	if err := db.DB.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": user})

}
func CreateUser(c *gin.Context) {
	var user models.User

	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON format"})
		return
	}

	if err := db.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusCreated, user)
}

func UpdateUser(c *gin.Context ){
	
	var user models.User

	//find user by id 

	idParam :=c.Param("id")
	id, err := strconv.Atoi(idParam)


	//check input
	if err != nil{
		c.JSON(http.StatusNotAcceptable, gin.H{"error": "Invalid user ID"})
		return
	}

	//find user
	if err := db.DB.Find(&user, id).Error; 
	err != nil{
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found!"})
		return
	}

	var updatedData models.User
	if err := c.ShouldBindJSON(&updatedData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	user.Name = updatedData.Name
	user.Email = updatedData.Email

	
	//save and return
	if err := db.DB.Save(&user).Error;
	err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
	}


	c.JSON(http.StatusOK, gin.H{"message": "User updated", "data": user})


}

func DeleteUser(c *gin.Context) {

	var user models.User

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

	if err := db.DB.Delete(&user, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "User Not Deleted"})
		return
	}

	c.Status(http.StatusNoContent)
}