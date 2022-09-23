package utils

import (
	"fmt"
	"ukiyo/config"
)

func AddStageToPath(path string) string {
	if config.ConfigSingleton.Stage != "" {
		return fmt.Sprintf("/%s%s", config.ConfigSingleton.Stage, path)
	}

	return path
}
