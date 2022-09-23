package config

import (
	"os"
)

type Config struct {
	Stage  string
	Server struct {
		Url string
	}
}

func (c *Config) Init() {
	if os.Getenv("STAGE") != "" {
		c.Stage = os.Getenv("STAGE")
	} else {
		c.Stage = ""
	}

	if os.Getenv("SERVER_URL") != "" {
		c.Server.Url = os.Getenv("SERVER_URL")
	} else {
		c.Server.Url = ":3000"
	}
}
