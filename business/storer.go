package business

import "ukiyo/models"

type Storer interface {
	GetPaintingsByArtist(artist string) []models.Painting
}
