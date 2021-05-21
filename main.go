package main

import (
	"gitlab.com/whizus/customer/pinto/cert-manager-webhook-pinto/pkg/dns"
	"os"

	"github.com/jetstack/cert-manager/pkg/acme/webhook/cmd"
)

// GroupName is the name under which the webhook will be available
var GroupName = os.Getenv("GROUP_NAME")

func main() {
	if GroupName == "" {
		panic("GROUP_NAME must be specified")
	}

	cmd.RunWebhookServer(GroupName,
		&dns.ProviderSolver{},
	)
}
