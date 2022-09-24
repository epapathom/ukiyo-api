package main

import (
	"ukiyo/server"
	"ukiyo/singleton"
)

func main() {
	singleton := new(singleton.Singleton)
	singleton.Init()

	server := new(server.Server)
	server.Init()
}
