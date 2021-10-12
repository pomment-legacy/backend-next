package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jessevdk/go-flags"
	"github.com/pomment/backend-next/server"
	"github.com/pomment/backend-next/server/config"
	"github.com/pomment/backend-next/server/dao"
	"github.com/pomment/backend-next/server/utils"
	"os"
	"path/filepath"
)

func main() {
	// required:"true"
	var opts struct {
		Config string `short:"d" long:"directory" description:"Data path"`
	}

	_, err := flags.ParseArgs(&opts, os.Args[1:])
	utils.FatalError(err)

	absPath, err := filepath.Abs(opts.Config)
	utils.FatalError(err)

	err = config.Read(absPath)
	utils.FatalError(err)

	dao.Init(opts.Config)
	engine := gin.Default()
	server.Init(engine, "v4")
	engine.Run(fmt.Sprintf("%s:%d", config.Content.APIHost, config.Content.APIPort))
}
