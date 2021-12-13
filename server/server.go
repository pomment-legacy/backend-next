package server

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/pomment/backend-next/server/config"
	"github.com/pomment/backend-next/server/controller"
	"github.com/pomment/backend-next/server/middleware"
)

type IncData struct {
	Variant string
}

func Init(engine *gin.Engine, prefix string) {
	r := engine

	// Sessions 管理
	var store = cookie.NewStore([]byte(config.Content.SiteAdmin.Salt))
	r.Use(sessions.Sessions("pomment_session", store))

	// 全局路由
	group := r.Group(prefix)
	{
		group.POST(("manage/login"), controller.Login)
		m := group.Group("manage", middleware.Verify)
		{
			m.GET("getThreads", controller.GetThreads)
			m.POST("getThread", controller.GetThread)
			m.POST("setThread", controller.SetThread)
			m.POST("getPosts", controller.GetPosts)
			m.POST("getPost", controller.GetPost)
			m.POST("setPost", controller.SetPost)
			m.POST("setSubPost", controller.SetSubPost)
		}
	}
}
