package main

import "long2ice/longurl/api"

func init() {
	app.Post("/", api.GenerateShortUrl)
	app.Get("/:path", api.VisitUrl)
}
