package cmd

import (
	"path/filepath"

	"github.com/stretchr/testify/assert"
)

func (ctx *Context) TestCreateProjectCmd() {
	assert := assert.New(ctx.T())

	rootCmd.SetArgs([]string{
		"--repodir", ctx.dir,
		"create-project", "testproject", "testgroup",
	})
	err := rootCmd.Execute()
	assert.Nil(err)

	expectedPaths := []string{
		"cluster-scope/components/project-admin-rolebindings/testgroup/kustomization.yaml",
		"cluster-scope/components/project-admin-rolebindings/testgroup/rbac.yaml",
		"cluster-scope/base/core/namespaces/testproject/kustomization.yaml",
		"cluster-scope/base/core/namespaces/testproject/namespace.yaml",
		"cluster-scope/base/user.openshift.io/groups/testgroup/kustomization.yaml",
		"cluster-scope/base/user.openshift.io/groups/testgroup/group.yaml",
	}

	for _, path := range expectedPaths {
		assert.FileExists(filepath.Join(ctx.dir, path))
	}

	// Should fail if namespace already exists
	rootCmd.SetArgs([]string{
		"--repodir", ctx.dir,
		"create-project", "testproject", "testgroup",
	})
	assert.PanicsWithError("Namespace testproject already exists", func() {
		rootCmd.Execute()
	})
}
