package dns

import (
	"fmt"
	"gitlab.com/whizus/gopinto"
	"strings"

	"github.com/jetstack/cert-manager/pkg/acme/webhook/apis/acme/v1alpha1"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

// ProviderSolver is the struct implementing the webhook.Solver interface
// for pinto DNS
type ProviderSolver struct {
	k8Client    kubernetes.Interface
	client      *gopinto.APIClient
	config      *Config
	apiKey      string
	provider    string
	environment string
}

func (p *ProviderSolver) Name() string {
	return p.getConfig().Name()
}

// Present is responsible for actually presenting the DNS record with the
// DNS provider.
// This method should tolerate being called multiple times with the same value.
// cert-manager itself will later perform a self check to ensure that the
// solver has correctly configured the DNS provider.
func (p *ProviderSolver) Present(ch *v1alpha1.ChallengeRequest) error {
	apiClient, err := p.getDomainAPIClient()
	if err != nil {
		return err
	}

	record, modelErr := p.getCreateRecordRequestModel(p.createRecordFromChallenge(ch))
	if modelErr != nil {
		return modelErr
	}
	_, _, creationErr := apiClient.RecordsApi.ApiDnsRecordsPost(p.config.getContext()).
		CreateRecordRequestModel(record).
		Execute()

	if creationErr.Error() != "" {
		return fmt.Errorf("failed to update DNS zone records: %w", creationErr)
	}

	return nil
}

// CleanUp should delete the relevant TXT record from the DNS provider console.
// If multiple TXT records exist with the same record name (e.g.
// _acme-challenge.example.com) then **only** the record with the same `key`
// value provided on the ChallengeRequest should be cleaned up.
// This is in order to facilitate multiple DNS validations for the same domain
// concurrently.
func (p *ProviderSolver) CleanUp(ch *v1alpha1.ChallengeRequest) error {
	apiClient, err := p.getDomainAPIClient()
	if err != nil {
		return err
	}

	//TODO BEGIN replace later when it is possible to delete by ID
	records, retrieveErr := p.getEntriesToPreserve(ch)
	if retrieveErr != nil {
		return retrieveErr
	}
	//TODO END

	_, deletionErr := apiClient.RecordsApi.ApiDnsRecordsDelete(p.getConfig().getContext()).
		Name(strings.TrimSuffix(strings.TrimSuffix(ch.ResolvedFQDN, ch.ResolvedZone), ".")).
		Zone(strings.TrimSuffix(ch.ResolvedZone, ".")).
		RecordType(gopinto.TXT).
		Provider(p.getConfig().Name()).
		Execute()

	if deletionErr.Error() != "" {
		return fmt.Errorf("failed to delete DNS zone records: %w", deletionErr)
	}

	//TODO BEGIN replace later when it is possible to delete by ID

	// re add entries
	for _, record := range records {

		recordModel, modelErr := p.getCreateRecordRequestModel(record)
		if modelErr != nil {
			return modelErr
		}

		_, _, creationErr := apiClient.RecordsApi.ApiDnsRecordsPost(p.getConfig().getContext()).
			CreateRecordRequestModel(recordModel).
			Execute()

		if creationErr.Error() != "" {
			return fmt.Errorf("failed to readd previous DNS zone records: %w", creationErr)
		}
	}
	//TODO END
	return nil
}

// Initialize will be called when the webhook first starts.
// This method can be used to instantiate the webhook, i.e. initialising
// connections or warming up caches.
// Typically, the kubeClientConfig parameter is used to build a Kubernetes
// k8Client that can be used to fetch resources from the Kubernetes API, e.g.
// Secret resources containing credentials used to authenticate with DNS
// provider accounts.
// The stopCh can be used to handle early termination of the webhook, in cases
// where a SIGTERM or similar signal is sent to the webhook process.
func (p *ProviderSolver) Initialize(kubeClientConfig *rest.Config, stopCh <-chan struct{}) error {

	cl, err := kubernetes.NewForConfig(kubeClientConfig)
	if err != nil {
		return fmt.Errorf("failed to get kubernetes k8Client: %w", err)
	}

	p.k8Client = cl

	client, domainErr := p.getDomainAPIClient()
	if domainErr != nil {
		return domainErr
	}

	p.client = client
	return nil
}
