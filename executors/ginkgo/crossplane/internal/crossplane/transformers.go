package crossplane

import (
	"fmt"

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

func RoleBindingSubjectTo(subject *Subject) manifest.Transformer {
	return func(u *unstructured.Unstructured) error {
		if u.GetKind() != "RoleBinding" {
			return nil
		}

		data, err := u.MarshalJSON()
		if err != nil {
			return err
		}

		value, err := subject.Marshal()
		if err != nil {
			return err
		}

		if data, err = jsonparser.Set(data, value, "subjects", "[0]"); err != nil {
			return err
		}

		return u.UnmarshalJSON(data)
	}
}

func ObjectProviderConfigRefTo(providerConfig string) manifest.Transformer {
	return func(u *unstructured.Unstructured) error {
		if u.GetKind() != "Object" {
			return nil
		}

		data, err := u.MarshalJSON()
		if err != nil {
			return err
		}

		value := []byte(fmt.Sprintf("%q", providerConfig))
		if data, err = jsonparser.Set(data, value, "spec", "providerConfigRef", "name"); err != nil {
			return err
		}

		return u.UnmarshalJSON(data)
	}
}

func ObjectManifestNamespaceTo(namespace string) manifest.Transformer {
	return func(u *unstructured.Unstructured) error {
		if u.GetKind() != "Object" {
			return nil
		}

		data, err := u.MarshalJSON()
		if err != nil {
			return err
		}

		value := []byte(fmt.Sprintf("%q", namespace))
		if data, err = jsonparser.Set(data, value, "spec", "forProvider", "manifest", "metadata", "namespace"); err != nil {
			return err
		}

		return u.UnmarshalJSON(data)
	}
}
