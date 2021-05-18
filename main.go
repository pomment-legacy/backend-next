package main

import (
	"github.com/jessevdk/go-flags"
	"github.com/pomment/backend-next/server"
	"os"
)

func main() {
	// required:"true"
	var opts struct {
		Config string `short:"c" long:"config" description:"Config file (in JSON)"`
	}

	_, err := flags.ParseArgs(&opts, os.Args[1:])
	if err != nil {
		os.Exit(1)
	}

	server.Start()
}
