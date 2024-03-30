package main

import "github.com/yaza-putu/golang-fiber-starter-mongo/internal/core"

func main() {
	// read env
	core.Env()
	// call mongo
	core.Mongo()
	// cal server
	core.Server()
}
