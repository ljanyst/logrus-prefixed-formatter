package main

import (
	"github.com/ljanyst/pre"
	"github.com/sirupsen/logrus"
)

var log = logrus.New()

func init() {
	formatter := new(pre.TextFormatter)
	formatter.MinPrefixWidth = 10
	log.Formatter = formatter
	log.Level = logrus.DebugLevel
}

func main() {
	defer func() {
		err := recover()
		if err != nil {
			// Fatal message
			log.WithFields(logrus.Fields{
				"omg":    true,
				"number": 100,
				"prefix": "main",
			}).Fatal("The ice breaks!")
		}
	}()

	// You could either provide a map key called `prefix` to add prefix
	log.WithFields(logrus.Fields{
		"prefix": "main",
		"animal": "walrus",
		"number": 8,
	}).Debug("Started observing beach")

	// Or you can simply add prefix in square brackets within message itself
	log.WithFields(logrus.Fields{
		"animal": "walrus",
		"size":   10,
		"prefix": "main",
	}).Debug("A group of walrus emerges from the ocean")

	// Warning message
	log.WithFields(logrus.Fields{
		"omg":    true,
		"number": 122,
	}).Warn("The group's number increased tremendously!")

	// Information message
	log.WithFields(logrus.Fields{
		"prefix":      "sensor",
		"temperature": -4,
	}).Info("Temperature changes")

	// Panic message
	log.WithFields(logrus.Fields{
		"prefix": "sensor",
		"animal": "orca",
		"size":   9009,
	}).Panic("It's over 9000!")
}
