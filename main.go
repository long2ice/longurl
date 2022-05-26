package main

import (
	"fmt"
	"github.com/long2ice/longurl/config"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.Fatal(app.Listen(fmt.Sprintf("%s:%s", config.ServerConfig.Host, config.ServerConfig.Port)))
}
