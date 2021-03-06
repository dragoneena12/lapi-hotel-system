package config

import (
	"log"
	"os"

	"gopkg.in/ini.v1"
)

type ConfigList struct {
	Debug     bool
	DbName    string
	SQLDriver string
	Port      int
	JWTSecret string
}

var Config ConfigList

func init() {
	cfg, err := ini.Load("config.ini")
	if err != nil {
		log.Printf("Failed to read file: %v", err)
		os.Exit(1)
	}

	Config = ConfigList{
		Debug:     cfg.Section("debug").Key("debug").MustBool(),
		DbName:    cfg.Section("db").Key("name").String(),
		SQLDriver: cfg.Section("db").Key("driver").String(),
		Port:      cfg.Section("web").Key("port").MustInt(),
		JWTSecret: cfg.Section("auth").Key("jwt_secret").String(),
	}
}
