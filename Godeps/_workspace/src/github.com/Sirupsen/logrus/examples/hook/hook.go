package main

import (
	"github.com/rochacon/docker-cron/Godeps/_workspace/src/github.com/Sirupsen/logrus"
	"github.com/rochacon/docker-cron/Godeps/_workspace/src/github.com/Sirupsen/logrus/hooks/airbrake"
)

var log = logrus.New()

func init() {
	log.Formatter = // default
		new(logrus.TextFormatter)
	log.Hooks.Add(airbrake.NewHook("https://example.com", "xyz", "development"))
}

func main() {
	log.WithFields(logrus.Fields{
		"animal": "walrus",
		"size":   10,
	}).Info("A group of walrus emerges from the ocean")

	log.WithFields(logrus.Fields{
		"omg":    true,
		"number": 122,
	}).Warn("The group's number increased tremendously!")

	log.WithFields(logrus.Fields{
		"omg":    true,
		"number": 100,
	}).Fatal("The ice breaks!")
}
