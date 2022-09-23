package business

import (
	"ukiyo/config"
	"ukiyo/models"
)

func GetPaintingByName(name string, language string) models.Painting {
	var storer Storer
	var painting models.Painting

	storer = config.DynamoDBSingleton

	if language == "japanese" {
		painting = storer.GetPaintingByName(name, "")
	} else if language == "english" {
		painting = storer.GetPaintingByName("", name)
	}

	return painting
}
