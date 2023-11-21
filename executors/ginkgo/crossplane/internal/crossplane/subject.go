package crossplane

import (
	"encoding/json"
)

func (r *Subject) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Subject struct {
	Kind      string `json:"kind"`
	Name      string `json:"name"`
	Namespace string `json:"namespace"`
}
