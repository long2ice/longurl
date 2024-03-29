package config

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"time"
)

type Server struct {
	Host          string `yaml:"host"`
	Port          string `yaml:"port"`
	LogTimezone   string `yaml:"logTimezone"`
	LogTimeFormat string `yaml:"logTimeFormat"`
}
type Url struct {
	Schema          string         `yaml:"schema"`
	Domain          string         `yaml:"domain"`
	Length          int            `yaml:"length"`
	AllowCustomPath bool           `yaml:"allowCustomPath"`
	ExpireSeconds   *time.Duration `yaml:"expireSeconds"`
	Unique          bool           `json:"unique"`
}
type Database struct {
	Type string `yaml:"type"`
	Dsn  string `yaml:"dsn"`
}
type Config struct {
	Server   *Server
	Url      *Url
	Database *Database
}

func init() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("config")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Fatal error config file: %v ", err)
	}
	var c Config
	err = viper.Unmarshal(&c)
	if err != nil {
		log.Fatalf("Unable to decode into struct, %v", err)
	}
	ServerConfig = c.Server
	DatabaseConfig = c.Database
	UrlConfig = c.Url
}

var (
	ServerConfig   *Server
	DatabaseConfig *Database
	UrlConfig      *Url
)
