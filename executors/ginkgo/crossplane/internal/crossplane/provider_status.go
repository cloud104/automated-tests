package crossplane

import (
	"encoding/json"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func UnmarshalProviderStatus(data []byte) (*ProviderStatus, error) {
	var r ProviderStatus
	err := json.Unmarshal(data, &r)

	return &r, err
}

func (r *ProviderStatus) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type ProviderStatus struct {
	Conditions        []metav1.Condition `json:"conditions"`
	CurrentIdentifier string             `json:"currentIdentifier"`
	CurrentRevision   string             `json:"currentRevision"`
}
