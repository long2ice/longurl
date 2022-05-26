package utils

import (
	"fmt"
	"github.com/long2ice/longurl/config"
)

var UrlConfig = config.UrlConfig

func FormatPath(path string) string {
	return fmt.Sprintf("%s://%s/%s", UrlConfig.Schema, UrlConfig.Domain, path)
}
