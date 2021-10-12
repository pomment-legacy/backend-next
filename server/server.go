package server

import (
	"github.com/gin-gonic/gin"
	"github.com/pomment/backend-next/server/controller"
)

type IncData struct {
	Variant string
}

func Init(engine *gin.Engine, prefix string) {
	r := engine
	group := r.Group(prefix)
	group.GET("manage/getThreads", controller.GetThreads)
	group.GET("manage/getThread", controller.GetThread)
	group.GET("test", controller.Test)
}
