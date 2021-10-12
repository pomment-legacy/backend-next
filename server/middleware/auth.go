package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/pomment/backend-next/server/utils"
)

func Verify(c *gin.Context) {
	token := c.GetHeader("Pomment-Token")
	if token == "123" {
		c.Next()
		return
	}
	c.JSON(403, utils.FailureRes(utils.MsgBadToken))
	c.Abort()
}
