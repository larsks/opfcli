package api

import (
	"os"
	"path/filepath"

	"github.com/operate-first/opfcli/constants"
	"github.com/stretchr/testify/require"
)

func (ctx *Context) TestCreateProject() {
	assert := require.New(ctx.T())

	err := ctx.api.CreateProject(
		"testproject",
		"testgroup",
		"test description",
		"",
		false,
	)
	assert.Nil(err)

	expectedPaths := []string{
		"cluster-scope/base/core/namespaces/testproject/namespace.yaml",
		"cluster-scope/base/core/namespaces/testproject/kustomization.yaml",
		"cluster-scope/base/user.openshift.io/groups/testgroup/group.yaml",
		"cluster-scope/base/user.openshift.io/groups/testgroup/kustomization.yaml",
		"cluster-scope/components/project-admin-rolebindings/testgroup/rbac.yaml",
	}

	compareWithExpected(assert, "testdata/CreateProject", ctx.dir, expectedPaths)
}

func (ctx *Context) TestCreateProjectQuota() {
	assert := require.New(ctx.T())

	// Should fail if quota does not exist
	err := ctx.api.CreateProject(
		"testproject",
		"testgroup",
		"test description",
		"testquota",
		false,
	)
	assert.EqualError(err, "quota testquota does not exist")

	// ---

	err = os.MkdirAll(filepath.Join(
		ctx.dir, ctx.api.AppName, constants.ComponentPath,
		"resourcequotas", "testquota",
	), 0755)
	assert.Nil(err)

	// Should succeed
	err = ctx.api.CreateProject(
		"testproject",
		"testgroup",
		"test description",
		"testquota",
		false,
	)
	assert.Nil(err)
}
