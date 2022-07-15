package api

import (
	"entgo.io/ent/dialect/sql"
	"github.com/gofiber/fiber/v2"
	"github.com/long2ice/longurl/config"
	"github.com/long2ice/longurl/db"
	"github.com/long2ice/longurl/ent"
	"github.com/long2ice/longurl/ent/url"
	"github.com/long2ice/longurl/sonyflake"
	"github.com/long2ice/longurl/utils"
	"github.com/mssola/user_agent"
	log "github.com/sirupsen/logrus"
	"time"
)

var UrlConfig = config.UrlConfig

type GenerateShortURL struct {
	URL      string    `json:"url" validate:"required" example:"https://github.com/long2ice/longurl"`
	ExpireAt time.Time `json:"expire_at"`
	Path     string    `json:"path" example:"longurl"`
	MaxTimes int       `json:"max_times" example:"1"`
}

func (g *GenerateShortURL) Handler(c *fiber.Ctx) error {
	if UrlConfig.Unique {
		fu, err := db.Client.Url.Query().Where(url.URL(g.URL), url.Or(url.ExpireAtGT(time.Now()), url.ExpireAtIsNil())).First(c.Context())
		if err != nil && !ent.IsNotFound(err) {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": err.Error(),
			})
		}
		if err == nil {
			u := fu.Update()
			if !g.ExpireAt.IsZero() {
				u.SetExpireAt(g.ExpireAt)
			}
			if g.Path != "" {
				u.SetPath(g.Path)
			}
			if g.MaxTimes != 0 {
				u.SetMaxTimes(g.MaxTimes)
			}
			u.SaveX(c.Context())
			return c.JSON(fiber.Map{
				"url": utils.FormatPath(fu.Path),
			})
		}
	}
	if g.Path == "" {
		id, err := sonyflake.SF.NextID()
		if err != nil {
			return err
		}
		str := utils.Encode(id)
		g.Path = str[len(str)-UrlConfig.Length:]
	} else {
		if !UrlConfig.AllowCustomPath {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "custom path is forbidden",
			})
		}
	}

	obj := db.Client.Url.Create().SetURL(g.URL).SetPath(g.Path).SetMaxTimes(g.MaxTimes)
	if g.ExpireAt.Unix() > 0 {
		if UrlConfig.ExpireSeconds != nil {
			if g.ExpireAt.Sub(time.Now()).Seconds() > float64(*UrlConfig.ExpireSeconds) {
				return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
					"message": "expire_at is larger than config value",
				})
			}
		}
		obj = obj.SetExpireAt(g.ExpireAt)
	} else {
		if UrlConfig.ExpireSeconds != nil {
			obj = obj.SetExpireAt(time.Now().Add(*UrlConfig.ExpireSeconds * time.Second))
		}
	}
	_, err := obj.Save(c.Context())
	if err != nil {
		return err
	}
	return c.JSON(fiber.Map{
		"url": utils.FormatPath(g.Path),
	})
}

type VisitURL struct {
	Path string `uri:"path"`
}

func (v *VisitURL) Handler(c *fiber.Ctx) error {
	u, err := db.Client.Url.Query().Where(
		url.Path(v.Path),
		url.Or(url.ExpireAtGT(time.Now()), url.ExpireAtIsNil()),
		url.Or(url.MaxTimesEQ(0), func(selector *sql.Selector) {
			selector.Where(sql.ColumnsGT(url.FieldMaxTimes, url.FieldCurrentTimes))
		}),
	).First(c.Context())
	if err != nil {
		return c.SendStatus(fiber.StatusNotFound)
	}
	if (u.ExpireAt != nil && u.ExpireAt.After(time.Now())) || u.ExpireAt == nil {
		v.createVisitLog(c, u)
		u.Update().AddCurrentTimes(1).SaveX(c.Context())
		return c.Redirect(u.URL)
	}
	return c.SendStatus(fiber.StatusNotFound)
}
func (v *VisitURL) createVisitLog(c *fiber.Ctx, url *ent.Url) {
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
