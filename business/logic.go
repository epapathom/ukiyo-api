package business

import (
	"ukiyo/adapters"
	"ukiyo/models"
)

type Logic struct {
	storer Storer
}

func (l *Logic) GetPaintingsByArtist(artist string) []models.Painting {
	var paintings []models.Painting

	l.storer = adapters.DynamoDBSingleton

	paintings = l.storer.GetPaintingsByArtist(artist)

	return paintings
}
