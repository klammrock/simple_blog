package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"simple_blog/blog_service"
	"simple_blog/config"
	"simple_blog/db"
	"simple_blog/level_log"

	"github.com/sirupsen/logrus"
	"github.com/weekface/mgorus"
)

const (
	// TODO: flag
	configFilePath = "./config.json"
	webPath        = "./web/"
	publicPath     = "./web/public/"
)

func getLogrusMongoDbHooker(config *config.Config) (logrus.Hook, error) {
	hooker, err := mgorus.NewHooker(fmt.Sprintf("%s:%d", config.Db.Host, config.Db.Port), config.Db.DbName, config.Db.LogCollection)
	if err != nil {
		return nil, err
	}

	return hooker, nil
}

func intiLogrus(hooker logrus.Hook) {
	logrus.AddHook(hooker)
	// disable output
	logrus.SetOutput(ioutil.Discard)
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetLevel(logrus.DebugLevel)
}

func logWithTag() *logrus.Entry {
	return logrus.WithField("Tag", "MAIN")
}

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

	// DB
	db, err := db.NewMongoDB(conf.Db)
	if err != nil {
		level_log.Fatalf("creating mongodb connection %v: %v", conf.Db, err)
	}
	level_log.Infof("mongodb connection created: %v", db)

	// Logrus MongoDB Hooker
	hooker, err := getLogrusMongoDbHooker(conf)
	if err != nil {
		level_log.Fatalf("init logrus mongo hooker: %v", err)
	}

	intiLogrus(hooker)
	level_log.Info("enabled logging to DB")

	blog_service.Do_work()

	// Service
	service, err := blog_service.New(config.IsDev(), conf.ServicePort, webPath, publicPath, db)
	if err != nil {
		level_log.Fatalf("creating service: %v", err)
	}
	level_log.Info("service created")

	level_log.Info("app initialized")
	level_log.Info("=============================================")

	logWithTag().Infof("starting service on port: %d", conf.ServicePort)
	if err := service.Start(); err != nil {
		logWithTag().Fatalf("start service: %v", err)
	}
}
