package crossplane

import (
	"github.com/totvs-cloud/go-manifest"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

func ByAPIVersion(apiVersion string) manifest.Filter {
	return func(u *unstructured.Unstructured) bool {
		return u.GetAPIVersion() == apiVersion
	}
}

func ByKind(kind string) manifest.Filter {
	return func(u *unstructured.Unstructured) bool {
		return u.GetKind() == kind
	}
}
