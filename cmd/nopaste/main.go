package main

import (
	"flag"

	"github.com/shogo82148/go-nopaste"
)

func main() {
	var config string
	flag.StringVar(&config, "c", "config.yaml", "path to config.yaml")
	flag.StringVar(&config, "config", "config.yaml", "path to config.yaml")
	flag.Parse()

	nopaste.Run(config)
}
