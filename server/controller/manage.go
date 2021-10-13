package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/pomment/backend-next/server/model"
	"github.com/pomment/backend-next/server/utils"
)

func GetThreads(c *gin.Context) {
	thread, err := model.GetThreads()
	if err != nil {
		c.JSON(500, utils.FailureRes(utils.MsgGeneralFailure))
		return
	}

	c.JSON(200, utils.SuccessRes(thread))
}

func GetThread(c *gin.Context) {
	url := c.Query("url")
	if url == "" {
		c.JSON(400, utils.FailureRes(utils.MsgBadArgument))
		return
	}

	thread, err := model.GetThread(url)
	if err != nil {
		c.JSON(500, utils.FailureRes(utils.MsgGeneralFailure))
		return
	}

	c.JSON(200, utils.SuccessRes(thread))
}

type SetThreadArgs struct {
	Url string `json:"url"`
	Title string `json:"title"`
}

func SetThread(c *gin.Context) {
	args := SetThreadArgs{}
	err := c.BindJSON(&args)
	if err != nil {
		c.JSON(400, utils.FailureRes(utils.MsgBadArgument))
		return
	}

	err = model.SetThread(args.Url, args.Title)
	if err != nil {
		c.JSON(500, utils.FailureRes(utils.MsgGeneralFailure))
		return
	}

	c.JSON(200, utils.SuccessRes(nil))
}

func GetPosts(c *gin.Context) {
	url := c.Query("url")
	if url == "" {
		c.JSON(400, utils.FailureRes(utils.MsgBadArgument))
		return
	}

	posts, err := model.GetPosts(url)
	if err != nil {
		c.JSON(500, utils.FailureRes(utils.MsgGeneralFailure))
		return
	}

	c.JSON(200, utils.SuccessRes(posts))
}