package logutils

import "github.com/jetstack/cert-manager/pkg/acme/webhook/apis/acme/v1alpha1"

func CreateChallengeFields(ch *v1alpha1.ChallengeRequest) map[string]interface{} {
	return map[string]interface{}{
		"zone":      ch.ResolvedZone,
		"fqdn":      ch.ResolvedFQDN,
		"namespace": ch.ResourceNamespace,
	}
}

func CreateModelFields(ch *v1alpha1.ChallengeRequest, model interface{}) map[string]interface{} {
	challengeFields := CreateChallengeFields(ch)
	result := challengeFields
	result["model"] = model
	return result
}

func CreateResponseFields(ch *v1alpha1.ChallengeRequest, response interface{}) map[string]interface{} {
	challengeFields := CreateChallengeFields(ch)
	result := challengeFields
	result["response"] = response
	return result
}
