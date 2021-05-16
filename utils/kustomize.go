package utils

import (
	"io/ioutil"
	"os/exec"
	"path/filepath"

	log "github.com/sirupsen/logrus"

	"github.com/operate-first/opfcli/models"
)

// WriteKustomization creates a kustomization.yaml adjacent to the given path.
func WriteKustomization(path string, resources []string, components []string) {
	kustom := models.NewKustomization()

	if len(resources) > 0 {
		kustom.Resources = resources
	}

	if len(components) > 0 {
		kustom.Components = components
	}

	kustomOut := models.ToYAML(kustom)

	err := ioutil.WriteFile(
		filepath.Join(filepath.Dir(path), "kustomization.yaml"),
		kustomOut, 0644,
	)
	if err != nil {
		log.Fatalf("failed to write kustomization: %v", err)
	}
}

// WriteComponent creates a Kustomization.yaml adjacent to the given path. This
// creates a Component rather than a Kustomization.
func WriteComponent(path string, resources []string) {
	kustom := models.NewKomponent()

	if len(resources) > 0 {
		kustom.Resources = resources
	}

	kustomOut := models.ToYAML(kustom)

	err := ioutil.WriteFile(
		filepath.Join(filepath.Dir(path), "kustomization.yaml"),
		[]byte(kustomOut), 0644,
	)
	if err != nil {
		log.Fatalf("failed to write kustomization: %v", err)
	}
}

// AddKustomizeComponent uses "kustomize edit add" to add a component
// to an existing kustomization file.
func AddKustomizeComponent(componentPath, kustomizePath string) {
	kustomize := exec.Command(
		"kustomize", "edit", "add", "component",
		componentPath,
	)
	kustomize.Dir = kustomizePath

	// NB: if the specified component does not exist, kustomize will fail to
	// edit the file and emit a log message but will not return an error code.
	if err := kustomize.Run(); err != nil {
		log.Fatalf("failed to edit kustomization: %v", err)
	}
}
