package cmd

import (
	"path/filepath"

	"github.com/stretchr/testify/assert"
)

func (ctx *Context) TestGrantAccessCmd() {
	var err error

	assert := assert.New(ctx.T())

	// Should fail if namespace doesn't exist.
	rootCmd.SetArgs([]string{
		"--repodir", ctx.dir,
		"grant-access", "testproject", "testgroup2", "admin",
	})
	assert.PanicsWithError("Namespace testproject does not exist", func() {
		rootCmd.Execute()
	})

	// ---

	rootCmd.SetArgs([]string{
		"--repodir", ctx.dir,
		"create-project", "testproject", "testgroup",
	})
	err = rootCmd.Execute()
	assert.Nil(err)

	// --

	// Should fail if target group doesn't exist
	rootCmd.SetArgs([]string{
		"--repodir", ctx.dir,
		"grant-access", "testproject", "testgroup2", "admin",
	})
	assert.PanicsWithError("Group testgroup2 does not exist", func() {
		rootCmd.Execute()
	})

	// ---

	rootCmd.SetArgs([]string{"--repodir", ctx.dir, "create-group", "testgroup2"})
	err = rootCmd.Execute()
	assert.Nil(err)

	// ---

	rootCmd.SetArgs([]string{
		"--repodir", ctx.dir,
		"grant-access", "testproject", "testgroup2", "admin",
	})
	err = rootCmd.Execute()
	assert.Nil(err)

	expectedPaths := []string{
		"cluster-scope/components/project-admin-rolebindings/testgroup/kustomization.yaml",
		"cluster-scope/components/project-admin-rolebindings/testgroup/rbac.yaml",
		"cluster-scope/components/project-admin-rolebindings/testgroup2/kustomization.yaml",
		"cluster-scope/components/project-admin-rolebindings/testgroup2/rbac.yaml",
		"cluster-scope/base/core/namespaces/testproject/kustomization.yaml",
		"cluster-scope/base/core/namespaces/testproject/namespace.yaml",
		"cluster-scope/base/user.openshift.io/groups/testgroup/kustomization.yaml",
		"cluster-scope/base/user.openshift.io/groups/testgroup/group.yaml",
		"cluster-scope/base/user.openshift.io/groups/testgroup2/kustomization.yaml",
		"cluster-scope/base/user.openshift.io/groups/testgroup2/group.yaml",
	}
	for _, path := range expectedPaths {
		assert.FileExists(filepath.Join(ctx.dir, path))
	}
}
