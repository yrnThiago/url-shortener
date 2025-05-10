package utils

import (
	"fmt"

	"github.com/yrnThiago/encurtador_url/config"
)

func GenerateShortUrl(id string) string {
	shortUrl := fmt.Sprintf("http://%s/%s", config.Env.ClientUrl, id)
	return shortUrl
}
