package main

import (
	"os"
	"simple_blog/blog_service"
	"simple_blog/config"
	"simple_blog/level_log"
	//"github.com/weekface/mgorus"
	//"gopkg.in/mgo.v2"
)

const (
	// TODO: flag
	configFilePath = "./config.json"
	webPath        = "./web/"
	publicPath     = "./web/public/"
)

func main() {
	level_log.Info("=============================================")
	level_log.Infof("starting app: %s", os.Args[0])

	// Config
	conf, err := config.LoadConfig(configFilePath)
	if err != nil {
		level_log.Fatalf("loading config file %s: %v", configFilePath, err)
	}

	if config.IsDev() {
		level_log.Infof("config file %s loaded: %v", configFilePath, conf)
	} else {
		// hide secrets in PROD mode
		level_log.Infof("config file %s loaded", configFilePath)
	}

	blog_service.Do_work()

	level_log.Info("app initialized")
	level_log.Info("=============================================")
}
