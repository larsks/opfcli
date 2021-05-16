package models

// Komponent represents a Kustomize Component. A Component is a collection of
// resources that can be included in a Kustomization file.
type Komponent struct {
	Resource  `yaml:",inline"`
	Resources []string `yaml:",omitempty"`
}

// Kustomization represents a kustomization file.
type Kustomization struct {
	Komponent  `yaml:",inline"`
	Components []string `yaml:",omitempty"`
}

// NewKustomization creates a new Kustomization object.
func NewKustomization() Kustomization {
	rsrc := Kustomization{
		Komponent: Komponent{
			Resource: Resource{
				APIVersion: "kustomize.config.k8s.io/v1beta1",
				Kind:       "Kustomization",
			},
		},
	}
	return rsrc
}

// NewKomponent creates a new Komponent object.
func NewKomponent() Komponent {
	rsrc := Komponent{
		Resource: Resource{
			APIVersion: "kustomize.config.k8s.io/v1alpha1",
			Kind:       "Component",
		},
	}
	return rsrc
}
