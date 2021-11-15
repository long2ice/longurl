package api

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/google/martian/log"
	"github.com/mssola/user_agent"
	"long2ice/longurl/config"
	"long2ice/longurl/db"
	"long2ice/longurl/ent"
	"long2ice/longurl/ent/url"
	"long2ice/longurl/schema"
	"long2ice/longurl/sonyflake"
	"long2ice/longurl/utils"
	"time"
)

func createVisitLog(c *fiber.Ctx, url *ent.Url) {
	userAgent := c.Request().Header.UserAgent()
	ua := user_agent.New(string(userAgent))
	engineName, engineVersion := ua.Engine()
	browserName, browserVersion := ua.Browser()
	ips := c.IPs()
	var ip string
	if len(ips) > 0 {
		ip = ips[0]
	} else {
		ip = c.IP()
	}
	_, err := db.Client.VisitLog.Create().
		SetURL(url).
		SetMobile(ua.Mobile()).
		SetBot(ua.Bot()).
		SetMozilla(ua.Mozilla()).
		SetPlatform(ua.Platform()).
		SetOs(ua.OS()).
		SetEngineName(engineName).
		SetEngineVersion(engineVersion).
		SetBrowserName(browserName).
		SetBrowserVersion(browserVersion).
		SetReferer(string(c.Request().Header.Referer())).
		SetIP(ip).
		Save(c.Context())
	if err != nil {
		log.Errorf("Create visit log error: %v", err)
	}
}
func VisitUrl(c *fiber.Ctx) error {
	path := c.Params("path")
	u, err := db.Client.Url.Query().Where(url.Path(path)).First(c.Context())
	if err != nil {
		return c.SendStatus(fiber.StatusNotFound)
	}
	if (u.ExpireAt != nil && u.ExpireAt.After(time.Now())) || u.ExpireAt == nil {
		createVisitLog(c, u)
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
	} else {
		if !UrlConfig.AllowCustomPath {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "custom path is not allowed",
			})
		}
	}

	obj := db.Client.Url.Create().SetURL(u.Url).SetPath(u.Path)
	if u.ExpireAt != nil {
		obj = obj.SetExpireAt(*u.ExpireAt)
	} else {
		obj = obj.SetExpireAt(time.Now().Add(UrlConfig.ExpireSeconds * time.Second))
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
