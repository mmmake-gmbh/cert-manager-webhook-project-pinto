package dns

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"gitlab.com/whizus/customer/pinto/cert-manager-webhook-pinto/internal/gopinto"
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
	configErr := p.getConfig().init(p.k8Client, ch)
	if configErr != nil {
		log.WithFields(map[string]interface{}{
			"zone":      ch.ResolvedZone,
			"fqdn":      ch.ResolvedFQDN,
			"namespace": ch.ResourceNamespace,
		}).WithError(configErr)
		return configErr
	}

	apiClient, err := p.getDomainAPIClient()
	if err != nil {
		log.WithFields(map[string]interface{}{
			"zone":      ch.ResolvedZone,
			"fqdn":      ch.ResolvedFQDN,
			"namespace": ch.ResourceNamespace,
		}).WithError(err)
		return err
	}

	record, modelErr := p.getCreateRecordRequestModel(p.createRecordFromChallenge(ch), ch.ResolvedZone)
	if modelErr != nil {
		log.WithFields(map[string]interface{}{
			"zone":      ch.ResolvedZone,
			"fqdn":      ch.ResolvedFQDN,
			"namespace": ch.ResourceNamespace,
		}).WithError(modelErr)
		return modelErr
	}
	requestModel := apiClient.RecordsApi.ApiDnsRecordsPost(p.config.getContext()).
		CreateRecordRequestModel(record)

	log.WithFields(map[string]interface{}{
		"model":     requestModel,
		"zone":      ch.ResolvedZone,
		"fqdn":      ch.ResolvedFQDN,
		"namespace": ch.ResourceNamespace,
	}).Trace("Prepared entry creation")
	_, response, creationErr := requestModel.Execute()

	if creationErr != nil {
		log.WithFields(map[string]interface{}{
			"zone":      ch.ResolvedZone,
			"fqdn":      ch.ResolvedFQDN,
			"namespace": ch.ResourceNamespace,
		}).WithError(modelErr)
		return creationErr
	}

	log.WithFields(map[string]interface{}{
		"response":  response,
		"zone":      ch.ResolvedZone,
		"fqdn":      ch.ResolvedFQDN,
		"namespace": ch.ResourceNamespace,
	}).Trace("Successfully created challenge")
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
		log.WithFields(map[string]interface{}{
			"zone":      ch.ResolvedZone,
			"fqdn":      ch.ResolvedFQDN,
			"namespace": ch.ResourceNamespace,
		}).WithError(err)
		return err
	}

	//TODO BEGIN replace later when it is possible to delete by ID
	records, retrieveErr := p.getEntriesToPreserve(ch)
	if retrieveErr != nil {
		log.WithFields(map[string]interface{}{
			"zone":      ch.ResolvedZone,
			"fqdn":      ch.ResolvedFQDN,
			"namespace": ch.ResourceNamespace,
		}).WithError(retrieveErr)
		return retrieveErr
	}

	log.WithFields(map[string]interface{}{
		"records":   records,
		"zone":      ch.ResolvedZone,
		"fqdn":      ch.ResolvedFQDN,
		"namespace": ch.ResourceNamespace,
	}).Trace("Retrieved list of TXT records to be readded")
	//TODO END

	deletionModel := apiClient.RecordsApi.ApiDnsRecordsDelete(p.getConfig().getContext()).
		Name(strings.TrimSuffix(strings.TrimSuffix(ch.ResolvedFQDN, ch.ResolvedZone), ".")).
		Zone(ch.ResolvedZone).
		Environment(p.getConfig().Environment()).
		RecordType(gopinto.TXT).
		Provider(p.getConfig().Name()).
		// if multiple entries with the same name are defined, we have to force the deletion of all
		RequestBody(map[string]string{
			"force": "true",
		})
	log.WithFields(map[string]interface{}{
		"model":     deletionModel,
		"zone":      ch.ResolvedZone,
		"fqdn":      ch.ResolvedFQDN,
		"namespace": ch.ResourceNamespace,
	}).Trace("Prepared deletion request")
	deletionResponse, deletionErr := deletionModel.Execute()

	if deletionErr != nil {
		log.WithFields(map[string]interface{}{
			"zone":      ch.ResolvedZone,
			"fqdn":      ch.ResolvedFQDN,
			"namespace": ch.ResourceNamespace,
		}).WithError(deletionErr)
		return deletionErr
	}
	log.WithFields(map[string]interface{}{
		"response":  deletionResponse,
		"zone":      ch.ResolvedZone,
		"fqdn":      ch.ResolvedFQDN,
		"namespace": ch.ResourceNamespace,
	}).Trace("Successfully deleted entries")

	//TODO BEGIN replace later when it is possible to delete by ID

	// re add entries
	for _, record := range records {

		recordModel, modelErr := p.getCreateRecordRequestModel(record, ch.ResolvedZone)
		if modelErr != nil {
			log.WithFields(map[string]interface{}{
				"zone":      ch.ResolvedZone,
				"fqdn":      ch.ResolvedFQDN,
				"namespace": ch.ResourceNamespace,
			}).WithError(modelErr)
			return modelErr
		}

		creationModel := apiClient.RecordsApi.ApiDnsRecordsPost(p.getConfig().getContext()).
			CreateRecordRequestModel(recordModel)
		log.WithFields(map[string]interface{}{
			"model":     creationModel,
			"zone":      ch.ResolvedZone,
			"fqdn":      ch.ResolvedFQDN,
			"namespace": ch.ResourceNamespace,
		}).Trace("Prepared recreation of re-add job")
		_, response, creationErr := creationModel.Execute()

		if creationErr != nil {
			log.WithFields(map[string]interface{}{
				"zone":      ch.ResolvedZone,
				"fqdn":      ch.ResolvedFQDN,
				"namespace": ch.ResourceNamespace,
			}).WithError(creationErr)
			return creationErr
		}

		log.WithFields(map[string]interface{}{
			"response":  response,
			"zone":      ch.ResolvedZone,
			"fqdn":      ch.ResolvedFQDN,
			"namespace": ch.ResourceNamespace,
		}).Trace("Successfully created re-add entry")
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
func (p *ProviderSolver) Initialize(kubeClientConfig *rest.Config, _ <-chan struct{}) error {

	cl, err := kubernetes.NewForConfig(kubeClientConfig)
	if err != nil {
		return fmt.Errorf("failed to get kubernetes k8Client: %w", err)
	}

	p.k8Client = cl

	return nil
}
