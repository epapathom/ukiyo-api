package configuration

import (
	"os"
)

type Configuration struct {
	Stage  string
	Server struct {
		Url string
	}
}

var ConfigurationSingleton *Configuration

func (c *Configuration) Init() {
	if os.Getenv("STAGE") != "" {
		c.Stage = os.Getenv("STAGE")
	} else {
		c.Stage = "dev"
	}

	if os.Getenv("SERVER_URL") != "" {
		c.Server.Url = os.Getenv("SERVER_URL")
	} else {
		c.Server.Url = ":3000"
	}
}
