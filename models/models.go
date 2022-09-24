package models

type Painting struct {
	Id           string `json:"id"`
	JapaneseName string `json:"japanese_name"`
	EnglishName  string `json:"english_name"`
	Artist       string `json:"artist"`
	Year         string `json:"year"`
}
