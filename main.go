package main

import (
	"github.com/jetstack/cert-manager/pkg/acme/webhook/cmd"
	log "github.com/sirupsen/logrus"
	"gitlab.com/whizus/customer/pinto/cert-manager-webhook-pinto/pkg/dns"
	"os"
)

// GroupName is the name under which the webhook will be available
var GroupName = os.Getenv("GROUP_NAME")

// allowed values are panic, fatal, error, warn, info, debug, trace
var logLevelEnv = os.Getenv("LOG_LEVEL")

func main() {
	// setup logging
	// Output to stdout instead of the default stderr
	log.SetOutput(os.Stdout)

	logLevel, logLevelErr := log.ParseLevel(logLevelEnv)
	if logLevelErr != nil {
		log.SetLevel(log.InfoLevel)
	} else {
		log.SetLevel(logLevel)
	}
	log.Info("Starting cert-manager-webhook-pinto")
	log.WithFields(map[string]interface{}{
		"logLevel":  log.GetLevel().String(),
		"groupName": GroupName,
	}).Debugf("application configuration used")

	if GroupName == "" {
		log.Panic("GROUP_NAME must be specified")
	}

	cmd.RunWebhookServer(GroupName,
		&dns.ProviderSolver{},
	)

	log.Info("Shutting down!")
	log.Exit(0)
}
