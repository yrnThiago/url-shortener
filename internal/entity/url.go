package entity

import (
	"github.com/google/uuid"
)

type Url struct {
	ID       string `bson:"_id"`
	FullUrl  string `bson:"full_url"`
	ShortUrl string `bson:"short_url"`
	Clicks   int    `bson:"clicks"`
}

func NewUrl(fullUrl string) *Url {
	id := uuid.New().String()

	return &Url{
		ID:      id,
		FullUrl: fullUrl,
		Clicks:  0,
	}
}

func (u *Url) SetShortUrl(shortUrl string) {
	u.ShortUrl = shortUrl
}
