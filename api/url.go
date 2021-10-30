package api

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"long2ice/longurl/config"
	"long2ice/longurl/db"
	"long2ice/longurl/ent/url"
	"long2ice/longurl/schema"
	"long2ice/longurl/sonyflake"
	"long2ice/longurl/utils"
	"time"
)

func VisitUrl(c *fiber.Ctx) error {
	path := c.Params("path")
	u, err := db.Client.Url.Query().Where(url.Path(path)).First(c.Context())
	if err != nil {
		return c.SendStatus(fiber.StatusNotFound)
	}
	if (u.ExpireAt != nil && u.ExpireAt.After(time.Now())) || u.ExpireAt == nil {
		return c.Redirect(u.URL)
	}
	return c.SendStatus(fiber.StatusNotFound)
}

var UrlConfig = config.UrlConfig

func GenerateShortUrl(c *fiber.Ctx) error {
	u := new(schema.Url)
	if err := c.BodyParser(u); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	if u.Path == "" {
		id, err := sonyflake.SF.NextID()
		if err != nil {
			return err
		}
		str := utils.Encode(id)
		u.Path = str[len(str)-UrlConfig.Length:]
	}

	obj := db.Client.Url.Create().SetURL(u.Url).SetPath(u.Path)
	if u.ExpireAt != nil {
		obj = obj.SetExpireAt(*u.ExpireAt)
	}
	_, err := obj.Save(c.Context())
	if err != nil {
		return err
	}
	shortUrl := fmt.Sprintf("%s://%s/%s", UrlConfig.Schema, UrlConfig.Domain, u.Path)
	return c.JSON(fiber.Map{
		"url": shortUrl,
	})
}
