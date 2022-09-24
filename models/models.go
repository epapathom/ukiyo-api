package models

type Painting struct {
	PaintingId   string `json:"painting_id"`
	JapaneseName string `json:"japanese_name"`
	EnglishName  string `json:"english_name"`
	Artist       string `json:"artist"`
	Year         string `json:"year"`
	S3Key        string `json:"s3_key"`
}
