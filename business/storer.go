package business

import "ukiyo/models"

type Storer interface {
	GetPaintingByName(japaneseName string, englishName string) models.Painting
}
