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
	group := r.Group(prefix)
	{
		m := group.Group("manage", middleware.Verify)
		{
			m.GET("getThreads", controller.GetThreads)
			m.GET("getThread", controller.GetThread)
			m.POST("setThread", controller.SetThread)
			m.GET("getPosts", controller.GetPosts)
			m.GET("getPost", controller.GetPost)
			m.POST("setPost", controller.SetPost)
			m.POST("setSubPost", controller.SetSubPost)
		}
	}
}
