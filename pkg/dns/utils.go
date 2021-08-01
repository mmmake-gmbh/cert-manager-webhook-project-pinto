package dns

import (
	"fmt"
	"github.com/camaoag/cert-manager-webhook-project-pinto/internal/gopinto"
	"github.com/jetstack/cert-manager/pkg/acme/webhook/apis/acme/v1alpha1"
	"github.com/jinzhu/copier"
	"github.com/sirupsen/logrus"
	cc "golang.org/x/oauth2/clientcredentials"
	"strings"
)

func (p *ProviderSolver) getDomainAPIClient() (*gopinto.APIClient, error) {
	config := gopinto.NewConfiguration()
	if config == nil {
		return nil, fmt.Errorf("failed to load config")
	}

	config.Servers[0].URL = p.getConfig().PintoApiURL()
	authClientConfig, err := configureOAuthClientConfig(p)
	if err != nil {
		return nil, err
	}
	config.HTTPClient = authClientConfig.Client(p.getConfig().getContext())

	return gopinto.NewAPIClient(config), nil
}

// transform gopinto.Record to gopinto.CreateRecordRequestModel
func (p *ProviderSolver) getCreateRecordRequestModel(record gopinto.Record, zone string) (gopinto.CreateRecordRequestModel, error) {
	var postRequestModel gopinto.CreateRecordRequestModel
	err := copier.Copy(&postRequestModel, &record)
	postRequestModel.SetZone(zone)

	return postRequestModel, err
}

// Returns a Record array in any case and optionally an error. Should an error occur the array will be partially filled or empty
func (p *ProviderSolver) getEntryList(ch *v1alpha1.ChallengeRequest) ([]gopinto.Record, error) {
	apiClient, err := p.getDomainAPIClient()
	if err != nil {
		return []gopinto.Record{}, err
	}

	apiOptions, createOptionsErr := p.createAPIOptions(nil)
	if createOptionsErr != nil {
		return nil, createOptionsErr
	}
	requestModel := apiClient.RecordApi.DnsApiRecordsGet(p.getConfig().getContext()).Zone(ch.ResolvedZone).
		Name(strings.TrimSuffix(strings.TrimSuffix(ch.ResolvedFQDN, ch.ResolvedZone), ".")).
		XApiOptions(apiOptions).
		RecordType(gopinto.RECORDTYPE_TXT)

	records, response, getError := requestModel.Execute()
	if getError != nil {
		logrus.Error(getError)
		return nil, getError
	}
	logrus.Trace(response)

	// TODO reimplement paging if pagination is fixed at the API. At the moment of writing entries are repeated on multiple pages
	//var aggregatedRecords []gopinto.Record
	//page := 0
	//for {
	//	records, _, getError := apiClient.RecordsApi.ApiDnsRecordsGet(p.getConfig().getContext()).
	//		Name(strings.TrimSuffix(strings.TrimSuffix(ch.ResolvedFQDN, ch.ResolvedZone), ".")).
	//		Zone(ch.ResolvedZone).
	//		Environment(p.getConfig().Environment()).
	//		RecordType(gopinto.TXT).
	//		Provider(p.getConfig().Name()).
	//		PageSize(pagingSize).
	//		PageToken(strconv.Itoa(page)).
	//		Execute()
	//
	//	// an error occurs return already gathered entries and the occurring error
	//	if getError.Error() != "" {
	//		logrus.Error(getError)
	//		return aggregatedRecords, getError
	//	}
	//	page++
	//
	//	aggregatedRecords = append(aggregatedRecords, records...)
	//	if len(records) != pagingSize {
	//		break
	//	}
	// }

	return records, nil
}

func (p *ProviderSolver) getEntriesToPreserve(ch *v1alpha1.ChallengeRequest) ([]gopinto.Record, error) {
	records, retrieveErr := p.getEntryList(ch)
	if retrieveErr != nil {
		return nil, retrieveErr
	}

	searchRecord := p.createRecordFromChallenge(ch)

	var foundRecords []gopinto.Record
	for _, record := range records {
		if hasSameNaming(record, searchRecord) {
			foundRecords = append(foundRecords, record)
		}
	}
	return foundRecords, nil
}

// return true if everything is the same except the data field. If Data is identical it will also return false
func hasSameNaming(a gopinto.Record, b gopinto.Record) bool {
	return a.Ttl == b.Ttl &&
		a.Name == b.Name &&
		a.Type == b.Type &&
		a.Class == b.Class &&
		a.Data != b.Data
}

func (p *ProviderSolver) createRecordFromChallenge(ch *v1alpha1.ChallengeRequest) gopinto.Record {
	ttl := int32(ttlDNS)
	return gopinto.Record{
		Name:  strings.TrimSuffix(strings.TrimSuffix(ch.ResolvedFQDN, ch.ResolvedZone), "."),
		Type:  gopinto.RECORDTYPE_TXT,
		Class: gopinto.RECORDCLASS_IN,
		Ttl:   &ttl,
		Data:  ch.Key,
	}
}

func (p *ProviderSolver) getConfig() *Config {
	if p.config == nil {
		p.config = &Config{}
	}
	return p.config
}

func configureOAuthClientConfig(p *ProviderSolver) (cc.Config, error) {
	tokenUrl := p.getConfig().OauthTokenURL()

	clientId := p.getConfig().OauthClientID()
	clientSecret := p.getConfig().OauthClientSecret()
	clientScope := p.getConfig().OauthClientScopes()

	oauthConfig := cc.Config{
		TokenURL:     tokenUrl,
		ClientID:     clientId,
		ClientSecret: clientSecret,
		Scopes:       clientScope,
	}
	return oauthConfig, nil
}

func (p ProviderSolver) createAPIOptions(meta map[string]string) (string, error) {
	apiOptions := new(gopinto.ApiOptions)
	accessOptions := new(gopinto.AccessOptions)
	accessOptions.SetEnvironment(p.getConfig().Environment())
	accessOptions.SetProvider(p.getConfig().Provider())
	accessOptions.SetCredentialsId(p.getConfig().CredentialsId())
	apiOptions.SetAccessOptions(*accessOptions)

	apiOptions.SetMeta(meta)

	marshalledJson, marshallErr := apiOptions.MarshalJSON()
	if marshallErr != nil {
		return "", marshallErr
	}
	result := string(marshalledJson)
	return result, nil
}
