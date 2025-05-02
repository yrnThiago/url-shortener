package server

import (
	"context"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/gofiber/fiber/v2/utils"

	"github.com/yrnThiago/encurtador_url/config"
)

type UrlInputDto struct {
	FullUrl string `json:"full_url"`
}

type UrlOutputDto struct {
	FullUrl  string `json:"full_url"`
	ShortUrl string `json:"short_url"`
	Clicks   int    `json:"clicks"`
}

func Init() {
	app := fiber.New()

	app.Use(requestid.New(requestid.Config{
		Next:       nil,
		Header:     fiber.HeaderXRequestID,
		Generator:  utils.UUID,
		ContextKey: "requestid",
	}))

	app.Post("/encurtaai", func(c *fiber.Ctx) error {
		var input UrlInputDto
		err := c.BodyParser(&input)
		if err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "bad request"})
		}

		newUrl := &config.Url{
			FullUrl:  input.FullUrl,
			ShortUrl: "short_url",
			Clicks:   0,
		}

		output := &UrlOutputDto{
			FullUrl:  newUrl.FullUrl,
			ShortUrl: newUrl.ShortUrl,
			Clicks:   newUrl.Clicks,
		}

		_, err = config.Conn.InsertOne(context.Background(), newUrl)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"message": fiber.StatusInternalServerError})
		}

		return c.Status(201).JSON(output)
	})

	log.Println("server listening on port " + "3000")

	log.Fatal(app.Listen(":" + "3000"))
}
