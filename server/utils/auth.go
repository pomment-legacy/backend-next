package utils

import "github.com/gin-gonic/gin"

func Verify(c *gin.Context) bool {
	token := c.GetHeader("Pomment-Token")
	if token == "123" {
		return true
	}
	return false
}
