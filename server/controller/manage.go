package controller

import (
	"log"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/pomment/backend-next/server/config"
	"github.com/pomment/backend-next/server/model"
	"github.com/pomment/backend-next/server/utils"
	"golang.org/x/crypto/bcrypt"
)

type LoginParam struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Login(c *gin.Context) {
	args := &LoginParam{}
	err := c.BindJSON(&args)
	if err != nil {
		c.JSON(400, utils.FailureRes(utils.MsgBadArgument))
		return
	}
	if args.Username != config.Content.SiteAdmin.Name {
		c.JSON(403, utils.FailureRes(utils.MsgBadLogin))
		return
	}
	err = bcrypt.CompareHashAndPassword([]byte(config.Content.SiteAdmin.Password), []byte(args.Password))
	if err != nil {
		c.JSON(403, utils.FailureRes(utils.MsgBadLogin))
		return
	}
	session := sessions.Default(c)
	session.Set("username", config.Content.SiteAdmin.Name)
	session.Set("salt", config.Content.SiteAdmin.Salt)
	session.Save()

	c.JSON(200, utils.SuccessRes(nil))
}

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
