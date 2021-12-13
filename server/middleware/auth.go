package middleware

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/pomment/backend-next/server/config"
	"github.com/pomment/backend-next/server/utils"
)

func Verify(c *gin.Context) {
	session := sessions.Default(c)
	username := session.Get("username")
	salt := session.Get("salt")
	if username == config.Content.SiteAdmin.Name && salt == config.Content.SiteAdmin.Salt {
		c.Next()
		return
	}
	c.JSON(403, utils.FailureRes(utils.MsgBadToken))
	c.Abort()
}
