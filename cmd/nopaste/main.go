package main

import (
	"github.com/shogo82148/go-nopaste"
)

func main() {
	nopaste.Run(&nopaste.Config{
		Root:    "/np",
		DataDir: "data",
	}, ":8080")
}
