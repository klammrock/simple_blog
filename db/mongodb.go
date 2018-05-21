package db

import (
	"fmt"

	"gopkg.in/mgo.v2"
)

type MongoDB struct {
	config Config
	events *mgo.Collection
	logs   *mgo.Collection
}

func NewMongoDB(config Config) (DBer, error) {
	if err := checkConfig(config); err != nil {
		return nil, err
	}

	url := fmt.Sprintf("%s:%d", config.Host, config.Port)

	session, err := mgo.Dial(url)
	if err != nil {
		return nil, err
	}

	events := session.DB(config.DbName).C(config.Collection)
	logs := session.DB(config.DbName).C(config.LogCollection)

	return &MongoDB{config, events, logs}, nil
}

func checkConfig(config Config) error {
	if len(config.Host) == 0 {
		return fmt.Errorf("Host can't be empty")
	}
	if config.Port <= 0 {
		return fmt.Errorf("Port <= 0")
	}
	if len(config.DbName) == 0 {
		return fmt.Errorf("DbName can't be empty")
	}
	if len(config.Collection) == 0 {
		return fmt.Errorf("Collection can't be empty")
	}
	if len(config.LogCollection) == 0 {
		return fmt.Errorf("LogCollection can't be empty")
	}

	return nil
}
