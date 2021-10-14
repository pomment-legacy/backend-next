package server

import (
	"github.com/gin-gonic/gin"
	"github.com/pomment/backend-next/server/controller"
	"github.com/pomment/backend-next/server/middleware"
)

type IncData struct {
	Variant string
}

func Init(engine *gin.Engine, prefix string) {
	r := engine
	group := r.Group(prefix, middleware.Verify)
	{
		group.GET("manage/getThreads", controller.GetThreads)
		group.GET("manage/getThread", controller.GetThread)
		group.POST("manage/setThread", controller.SetThread)
		group.GET("manage/getPosts", controller.GetPosts)
		group.GET("manage/getPost", controller.GetPost)
		group.POST("manage/setPost", controller.SetPost)
	}
}
