package crossplane

import (
	"github.com/buger/jsonparser"
	"github.com/totvs-cloud/go-manifest"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

func NameTo(name string) manifest.Transformer {
	return func(u *unstructured.Unstructured) error {
		u.SetName(name)
		return nil
	}
}

func NamespaceTo(namespace string) manifest.Transformer {
	return func(u *unstructured.Unstructured) error {
		u.SetNamespace(namespace)
		return nil
	}
}

func RoleBindingSubjectTo(serviceAccount string) manifest.Transformer {
	return func(u *unstructured.Unstructured) error {
		if u.GetKind() != "RoleBinding" {
			return nil
		}

		data, err := u.MarshalJSON()
		if err != nil {
			return err
		}

		if data, err = jsonparser.Set(data, []byte(serviceAccount), "subjects", "[0]", "name"); err != nil {
			return err
		}

		return u.UnmarshalJSON(data)
	}
}
