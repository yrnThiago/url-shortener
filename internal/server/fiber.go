package server

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.uber.org/zap"

	"github.com/yrnThiago/encurtador_url/config"
	"github.com/yrnThiago/encurtador_url/internal/entity"
	"github.com/yrnThiago/encurtador_url/internal/utils"
)

type UrlInputDto struct {
	FullUrl string `json:"full_url" validate:"http_url"`
}

type UrlOutputDto struct {
	FullUrl  string `json:"full_url"`
	ShortUrl string `json:"short_url"`
	Clicks   int    `json:"clicks"`
}

const ApiEndpoint = "/encurtaai"

func Init() {
	app := fiber.New()
	app.Use(cors.New())
	api := app.Group(ApiEndpoint)

	api.Post("/", func(c *fiber.Ctx) error {
		var input UrlInputDto
		err := c.BodyParser(&input)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "bad request"})
		}

		err = utils.ValidateStruct(input)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).
				JSON(fiber.Map{"error": "url needs to be valid http url"})
		}

		newUrl := entity.NewUrl(input.FullUrl)
		newUrl.SetShortUrl(utils.GenerateShortUrl(newUrl.ID))

		output := &UrlOutputDto{
			FullUrl:  newUrl.FullUrl,
			ShortUrl: newUrl.ShortUrl,
			Clicks:   newUrl.Clicks,
		}

		_, err = config.Conn.InsertOne(context.Background(), newUrl)
		if err != nil {
			config.Logger.Warn(
				"failed to short new url",
				zap.Error(err),
			)

			return c.Status(fiber.StatusInternalServerError).
				JSON(fiber.Map{"message": fiber.StatusInternalServerError})
		}

		config.Logger.Info(
			"new short url",
			zap.String("short_url", output.ShortUrl),
		)

		return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": output.ShortUrl})
	})

	api.Get("/:id", func(c *fiber.Ctx) error {
		var shortUrl entity.Url
		id := c.Params("id")
		if id == ":id" {
			return c.Status(fiber.StatusBadRequest).
				JSON(fiber.Map{"error": fiber.ErrBadRequest.Message})
		}

		err := config.Conn.FindOne(context.Background(), bson.M{"_id": id}).Decode(&shortUrl)
		if err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "url not found"})
		}

		clicksUpdated := shortUrl.Clicks + 1
		updateUrl := bson.M{"$set": bson.M{"clicks": clicksUpdated}}

		_, err = config.Conn.UpdateOne(context.Background(), bson.M{"_id": id}, updateUrl)
		if err != nil {
			config.Logger.Warn(
				"failed to update url",
				zap.Error(err),
			)

			return c.Status(fiber.StatusInternalServerError).
				JSON(fiber.Map{"error": fiber.ErrInternalServerError.Message})
		}

		config.Logger.Info(
			"someone clicked on",
			zap.String("url", shortUrl.ShortUrl),
			zap.Int("clicks", clicksUpdated),
		)

		return c.Redirect(shortUrl.FullUrl)
	})

	config.Logger.Info(
		"server listening",
		zap.String("port", config.Env.Port),
	)

	config.Logger.Fatal(
		"something went wrong",
		zap.Error(app.Listen(":"+config.Env.Port)),
	)
}
