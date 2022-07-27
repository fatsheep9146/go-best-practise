package main

import (
	"github.com/sirupsen/logrus"
	prefixed "github.com/x-cray/logrus-prefixed-formatter"
)

var log = logrus.New()

func init() {
	log.Formatter = new(prefixed.TextFormatter)
	log.Level = logrus.DebugLevel
}

func main() {
	log.WithFields(logrus.Fields{
		"prefix": "main",
		"animal": "walrus",
		"number": 8,
	}).Debug("Started observing beach")

	log.WithFields(logrus.Fields{
		"prefix":      "sensor",
		"temperature": -4,
	}).Info("Temperature changes")
}
