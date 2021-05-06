package dns

import (
	"fmt"
	"github.com/jetstack/cert-manager/pkg/acme/webhook/apis/acme/v1alpha1"
	"github.com/jinzhu/copier"
	"gitlab.com/whizus/gopinto"
	"strconv"
	"strings"
)

const pagingSize = 20
const ttlDNS = 60

func (p *ProviderSolver) getDomainAPI() (*gopinto.APIClient, error) {
	config := gopinto.NewConfiguration()
	if config == nil {
		return nil, fmt.Errorf("failed to load config")
	}
	client := gopinto.NewAPIClient(config)

	// TODO refactor

	return client, nil
}

// transform gopinto.Record to gopinto.CreateRecordRequestModel
func (p *ProviderSolver) getCreateRecordRequestModel(record gopinto.Record) (gopinto.CreateRecordRequestModel, error) {
	var postRequestModel gopinto.CreateRecordRequestModel
	err := copier.Copy(&record, &postRequestModel)
	return postRequestModel, err
}

// Returns a Record array in any case and optionally an error. Should an error occur the array will be partially filled or empty
func (p *ProviderSolver) getEntryList(ch *v1alpha1.ChallengeRequest) ([]gopinto.Record, error) {
	apiClient, err := p.getDomainAPI()
	if err != nil {
		return []gopinto.Record{}, err
	}

	var aggregatedRecords []gopinto.Record
	page := 0
	for {
		records, _, getError := apiClient.RecordsApi.ApiDnsRecordsGet(p.getContext()).
			Name(strings.TrimSuffix(strings.TrimSuffix(ch.ResolvedFQDN, ch.ResolvedZone), ".")).
			Zone(strings.TrimSuffix(ch.ResolvedZone, ".")).
			RecordType(gopinto.TXT).
			Provider(p.Name()).
			PageSize(pagingSize).
			PageToken(strconv.Itoa(page)).
			Execute()

		// an error occurs return already gathered entries and the occurring error
		if getError.Error() != "" {
			return aggregatedRecords, getError
		}

		aggregatedRecords = append(aggregatedRecords, records...)
		if len(records) != pagingSize {
			break
		}
	}
	return aggregatedRecords, nil
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
	ttl := int64(ttlDNS)
	return gopinto.Record{
		Name:  strings.TrimSuffix(strings.TrimSuffix(ch.ResolvedFQDN, ch.ResolvedZone), "."),
		Type:  gopinto.TXT,
		Class: p.Name(),
		Ttl:   &ttl,
		Data:  strconv.Quote(ch.Key),
	}
}
