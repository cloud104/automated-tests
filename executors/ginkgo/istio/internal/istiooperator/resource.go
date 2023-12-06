package istiooperator

import (
	"encoding/json"
)

func UnmarshalResource(data []byte) (*Resource, error) {
	var r Resource
	err := json.Unmarshal(data, &r)

	return &r, err
}

func (r *Resource) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Resource struct {
	APIVersion string         `json:"apiVersion"`
	Kind       string         `json:"kind"`
	Metadata   Metadata       `json:"metadata"`
	Status     ResourceStatus `json:"status"`
}

type Metadata struct {
	Name      string `json:"name"`
	Namespace string `json:"namespace"`
}

type ResourceStatus struct {
	ComponentStatus Components `json:"componentStatus"`
	Status          string     `json:"status"`
}

type Components struct {
	Base            ComponentStatus `json:"Base"`
	EgressGateways  ComponentStatus `json:"EgressGateways"`
	IngressGateways ComponentStatus `json:"IngressGateways"`
	Pilot           ComponentStatus `json:"Pilot"`
}

type ComponentStatus struct {
	Status string `json:"status"`
}
