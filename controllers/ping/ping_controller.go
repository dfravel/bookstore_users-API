package ping

import "github.com/gin-gonic/gin"

// Ping to make sure the server is up and running
func Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
