package dns

import (
	"fmt"
	"github.com/jetstack/cert-manager/pkg/acme/webhook/apis/acme/v1alpha1"
	"gitlab.com/whizus/gopinto"
	"strconv"
	"strings"
)

func (p *ProviderSolver) getDomainAPI(ch *v1alpha1.ChallengeRequest) (*gopinto.APIClient, error) {
	config := gopinto.NewConfiguration()
	if config == nil {
		return nil, fmt.Errorf("failed to load config")
	}
	client := gopinto.NewAPIClient(config)

	// TODO refactor

	return client, nil
}

func (p *ProviderSolver) getCreateRecordRequestModel(ch *v1alpha1.ChallengeRequest) gopinto.CreateRecordRequestModel {
	ttl := int32(60)
	postRequestModel := gopinto.CreateRecordRequestModel{
		Provider: p.Name(),
		Zone:     strings.TrimSuffix(ch.ResolvedZone, "."),
		Name:     strings.TrimSuffix(strings.TrimSuffix(ch.ResolvedFQDN, ch.ResolvedZone), "."),
		Type:     gopinto.TXT,
		Data:     strconv.Quote(ch.Key),
		Ttl:      &ttl,
	}
	return postRequestModel
}
