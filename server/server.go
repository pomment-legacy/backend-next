package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type IncData struct {
	Variant string
}

func Start() {
	r := gin.Default()
	// placeholder
	r.GET("/pomment/:type/:method", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"type":   c.Param("type"),
			"method": c.Param("method"),
		})
	})
	r.Run(fmt.Sprintf("127.0.0.1:%d", 8080)) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
