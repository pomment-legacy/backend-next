package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/pomment/backend-next/server/model"
	"github.com/pomment/backend-next/server/utils"
)

func GetThreads(c *gin.Context) {
	if !utils.Verify(c) {
		c.JSON(403, utils.FailureRes(utils.MsgBadToken))
		return
	}

	thread, err := model.GetThreads()
	if err != nil {
		c.JSON(500, utils.FailureRes(utils.MsgGeneralFailure))
		return
	}

	c.JSON(200, utils.SuccessRes(thread))
}

func GetThread(c *gin.Context) {
	if !utils.Verify(c) {
		c.JSON(403, utils.FailureRes(utils.MsgBadToken))
		return
	}

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
