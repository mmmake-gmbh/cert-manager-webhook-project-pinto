package main

import (
	"github.com/sirupsen/logrus"
	"os"
)

var log = logrus.New()
// GroupName is the name under which the webhook will be available
var GroupName = os.Getenv("GROUP_NAME")

func main() {
	if GroupName == "" {
		log.Errorf("Environment variable %s is required for program to start! Shutting down.", "GROUP_NAME")
		os.Exit(1)
	}
	// TODO: implement me
}
