package main

import (
	"ukiyo/config"
	"ukiyo/server"
)

func main() {
	singleton := new(config.Singleton)
	singleton.Init()

	server := new(server.Server)
	server.Init()
}
