package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/operate-first/opfcli/models"
	"github.com/operate-first/opfcli/utils"
	log "github.com/sirupsen/logrus"
)

func createNamespace(projectName, projectOwner, projectDescription string) {
	appName := config.GetString("app-name")
	path := filepath.Join(repoDirectory, appName, namespacePath, projectName, "namespace.yaml")

	if utils.PathExists(filepath.Dir(path)) {
		log.Fatalf("namespace %s already exists", projectName)
	}

	ns := models.NewNamespace(projectName, projectOwner, projectDescription)
	nsOut := models.ToYAML(ns)

	log.Printf("writing namespace definition to %s", filepath.Dir(path))
	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		log.Fatalf("failed to create namespace directory: %v", err)
	}

	err := ioutil.WriteFile(path, nsOut, 0644)
	if err != nil {
		log.Fatalf("failed to write namespace file: %v", err)
	}

	utils.WriteKustomization(
		path,
		[]string{"namespace.yaml"},
		[]string{
			filepath.Join(componentRelPath, "project-admin-rolebindings", projectOwner),
		},
	)
}

func createRoleBinding(projectName, groupName, roleName string) {
	appName := config.GetString("app-name")
	bindingName := fmt.Sprintf("project-%s-rolebindings", roleName)
	path := filepath.Join(
		repoDirectory, appName, componentPath,
		bindingName, groupName, "rbac.yaml",
	)

	if utils.PathExists(filepath.Dir(path)) {
		log.Printf("rolebinding already exists (continuing)")
		return
	}

	rbac := models.NewRoleBinding(
		fmt.Sprintf("namespace-%s-%s", roleName, groupName),
		roleName,
	)
	rbac.AddGroup(groupName)
	rbacOut := models.ToYAML(rbac)

	log.Printf("writing rbac definition to %s", filepath.Dir(path))
	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		log.Fatalf("failed to create rolebinding directory: %v", err)
	}

	err := ioutil.WriteFile(path, rbacOut, 0644)
	if err != nil {
		log.Fatalf("failed to write rbac: %v", err)
	}

	utils.WriteComponent(
		path,
		[]string{"rbac.yaml"},
	)
}

func createAdminRoleBinding(projectName, projectOwner string) {
	createRoleBinding(projectName, projectOwner, "admin")
}

func createGroup(projectOwner string) {
	appName := config.GetString("app-name")
	path := filepath.Join(repoDirectory, appName, groupPath, projectOwner, "group.yaml")

	if utils.PathExists(filepath.Dir(path)) {
		log.Printf("group already exists (continuing)")
		return
	}

	group := models.NewGroup(projectOwner)
	groupOut := models.ToYAML(group)

	log.Printf("writing group definition to %s", filepath.Dir(path))
	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		log.Fatalf("failed to create group directory: %v", err)
	}

	err := ioutil.WriteFile(path, groupOut, 0644)
	if err != nil {
		log.Fatalf("failed to write group: %v", err)
	}

	utils.WriteKustomization(
		path,
		[]string{"group.yaml"},
		nil,
	)
}

func addGroupRBAC(projectName, groupName, roleName string) {
	appName := config.GetString("app-name")
	bindingName := fmt.Sprintf("project-%s-rolebindings", roleName)

	nsPath := filepath.Join(
		repoDirectory, appName, namespacePath, projectName,
	)

	groupPath := filepath.Join(
		repoDirectory, appName, groupPath, groupName,
	)

	if !utils.PathExists(nsPath) {
		log.Fatalf("namespace %s does not exist", projectName)
	}

	if !utils.PathExists(groupPath) {
		log.Fatalf("group %s does not exist", groupName)
	}

	createRoleBinding(projectName, groupName, roleName)

	log.Printf("granting %s role %s on %s", groupName, roleName, projectName)
	utils.AddKustomizeComponent(
		filepath.Join(componentRelPath, bindingName, groupName),
		nsPath,
	)
}
