package main

import (
	"fmt"
	"github.com/jessevdk/go-flags"
	"github.com/pomment/backend-next/server/config"
	"os"
)

func main() {
	// required:"true"
	var opts struct {
		Config string `short:"d" long:"directory" description:"Data path"`
	}

	_, err := flags.ParseArgs(&opts, os.Args[1:])
	if err != nil {
		os.Exit(1)
	}

	err = config.Read(opts.Config)
	if err != nil {
		os.Exit(1)
	}
	fmt.Printf("APIHost 是 %s\n", config.Content.APIHost)
	fmt.Printf("APIPort 是 %d\n", config.Content.APIPort)
}
