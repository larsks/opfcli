package cmd

import (
	"path/filepath"

	"github.com/stretchr/testify/assert"
)

func (ctx *Context) TestCreateGroupCmd() {
	assert := assert.New(ctx.T())

	rootCmd.SetArgs([]string{"--repodir", ctx.dir, "create-group", "testgroup"})
	err := rootCmd.Execute()
	assert.Nil(err)

	expectedPaths := []string{
		"cluster-scope/base/user.openshift.io/groups/testgroup/group.yaml",
		"cluster-scope/base/user.openshift.io/groups/testgroup/kustomization.yaml",
	}

	for _, path := range expectedPaths {
		assert.FileExists(filepath.Join(ctx.dir, path))
	}
}
