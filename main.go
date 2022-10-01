package main

import (
	"ukiyo/server"
	"ukiyo/singleton"
)

// @title           Ukiyo API
// @version         1.0
// @description     An API to store and retrieve Ukiyo-e paintings.

// @host      localhost:3000
// @BasePath  /dev
func main() {
	singleton := new(singleton.Singleton)
	singleton.Init()

	server := new(server.Server)
	server.Init()
}
