package main

import (
	"image-processor/server"
)

func main() {
	server.NewHttpServer().Run()
}
