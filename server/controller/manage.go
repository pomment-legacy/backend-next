package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/pomment/backend-next/server/model"
	"github.com/pomment/backend-next/server/utils"
	"log"
)

func GetThreads(c *gin.Context) {
	thread, err := model.GetThreads()
	if err != nil {
		log.Printf("Error: %s", err)
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
		log.Printf("Error: %s", err)
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
		log.Printf("Error: %s", err)
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
		log.Printf("Error: %s", err)
		c.JSON(500, utils.FailureRes(utils.MsgGeneralFailure))
		return
	}

	c.JSON(200, utils.SuccessRes(posts))
}

func GetPost(c *gin.Context) {
	url := c.Query("url")
	uuid := c.Query("uuid")
	if url == "" || uuid == "" {
		c.JSON(400, utils.FailureRes(utils.MsgBadArgument))
		return
	}

	post, _, err := model.GetPost(url, uuid)
	if err != nil {
		log.Printf("Error: %s", err)
		c.JSON(500, utils.FailureRes(utils.MsgGeneralFailure))
		return
	}

	c.JSON(200, utils.SuccessRes(post))
}

func SetPost(c *gin.Context) {
	args := model.SetPostParam{}
	err := c.BindJSON(&args)
	if err != nil || args.Url == "" || args.UUID == "" {
		c.JSON(400, utils.FailureRes(utils.MsgBadArgument))
		return
	}

	data, err := model.SetPost(args)
	if err != nil {
		log.Printf("Error: %s", err)
		c.JSON(500, utils.FailureRes(utils.MsgGeneralFailure))
		return
	}

	c.JSON(200, utils.SuccessRes(data))
}

func SetSubPost(c *gin.Context) {
	args := model.SetSubPostParam{}
	err := c.BindJSON(&args)
	if err != nil || args.Url == "" || args.UUID == "" || args.ParentUUID == "" {
		c.JSON(400, utils.FailureRes(utils.MsgBadArgument))
		return
	}

	data, err := model.SetSubPost(args)
	if err != nil {
		log.Printf("Error: %s", err)
		c.JSON(500, utils.FailureRes(utils.MsgGeneralFailure))
		return
	}

	c.JSON(200, utils.SuccessRes(data))
}
