package main

import (
	"github.com/long2ice/fibers/router"
	"github.com/long2ice/longurl/api"
)

var (
	generateURL = router.New(
		&api.GenerateShortURL{},
		router.Summary("Generate short url"),
	)
	visitURL = router.New(
		&api.VisitURL{},
		router.Summary("Visit short url"),
	)
)

func init() {
	app.Post("/", generateURL)
	app.Get("/:path", visitURL)
}
