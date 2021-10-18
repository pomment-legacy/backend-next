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

type GetThreadParam struct {
	Url string `json:"url"`
}

func GetThread(c *gin.Context) {
	data := GetThreadParam{}
	err := c.BindJSON(&data)
	if err != nil || data.Url == "" {
		c.JSON(400, utils.FailureRes(utils.MsgBadArgument))
		return
	}

	thread, err := model.GetThread(data.Url)
	if err != nil {
		log.Printf("Error: %s", err)
		c.JSON(500, utils.FailureRes(utils.MsgGeneralFailure))
		return
	}

	c.JSON(200, utils.SuccessRes(thread))
}

type SetThreadParam struct {
	Url   string `json:"url"`
	Title string `json:"title"`
}

func SetThread(c *gin.Context) {
	args := SetThreadParam{}
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

type GetPostsParam struct {
	Url string `json:"url"`
}

func GetPosts(c *gin.Context) {
	data := GetPostsParam{}
	err := c.BindJSON(&data)
	if err != nil || data.Url == "" {
		c.JSON(400, utils.FailureRes(utils.MsgBadArgument))
		return
	}

	posts, err := model.GetPosts(data.Url)
	if err != nil {
		log.Printf("Error: %s", err)
		c.JSON(500, utils.FailureRes(utils.MsgGeneralFailure))
		return
	}

	c.JSON(200, utils.SuccessRes(posts))
}

type GetPostParam struct {
	Url  string `json:"url"`
	UUID string `json:"uuid"`
}

func GetPost(c *gin.Context) {
	data := GetPostParam{}
	err := c.BindJSON(&data)
	if err != nil || data.Url == "" || data.UUID == "" {
		c.JSON(400, utils.FailureRes(utils.MsgBadArgument))
		return
	}

	post, _, err := model.GetPost(data.Url, data.UUID)
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
