package user

import "github.com/gin-gonic/gin"

// GetUser .
func GetUser(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "you're calling the get user function",
	})
}

// GetUserByID .
func GetUserByID(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "you're calling the get user bu id function",
	})
}
