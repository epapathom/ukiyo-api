package business

import (
	"ukiyo/config"
	"ukiyo/models"
)

func GetPaintingsByArtist(artist string) []models.Painting {
	var storer Storer
	var paintings []models.Painting

	storer = config.DynamoDBSingleton

	paintings = storer.GetPaintingsByArtist(artist)

	return paintings
}
