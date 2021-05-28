package api

import (
	"path/filepath"

	"github.com/stretchr/testify/require"
)

func (ctx *Context) TestAddGroupRBAC() {
	assert := require.New(ctx.T())

	// Should fail if project does not exist
	err := ctx.api.AddGroupRBAC(
		"testproject",
		"testgroup2",
		"admin",
	)
	assert.EqualError(err, "namespace testproject does not exist")

	// ---

	// Should fail if group does not exist
	err = ctx.api.CreateProject(
		"testproject",
		"testgroup",
		"test description",
		"",
		false,
	)
	assert.Nil(err)

	err = ctx.api.AddGroupRBAC(
		"testproject",
		"testgroup2",
		"admin",
	)
	assert.EqualError(err, "group testgroup2 does not exist")

	// ---

	// Should work if both project and group exist
	err = ctx.api.CreateGroup(
		"testgroup2",
		false,
	)
	assert.Nil(err)
	assert.FileExists(filepath.Join(
		ctx.dir,
		"cluster-scope/base/user.openshift.io/groups/testgroup2/group.yaml",
	))

	err = ctx.api.AddGroupRBAC(
		"testproject",
		"testgroup2",
		"admin",
	)
	assert.Nil(err)

	expectedPaths := []string{
		"cluster-scope/base/core/namespaces/testproject/namespace.yaml",
		"cluster-scope/base/core/namespaces/testproject/kustomization.yaml",
		"cluster-scope/base/user.openshift.io/groups/testgroup/group.yaml",
		"cluster-scope/base/user.openshift.io/groups/testgroup/kustomization.yaml",
		"cluster-scope/base/user.openshift.io/groups/testgroup2/group.yaml",
		"cluster-scope/base/user.openshift.io/groups/testgroup2/kustomization.yaml",
		"cluster-scope/components/project-admin-rolebindings/testgroup/rbac.yaml",
		"cluster-scope/components/project-admin-rolebindings/testgroup2/rbac.yaml",
	}

	compareWithExpected(assert, "testdata/AddGroupRBAC", ctx.dir, expectedPaths)
}
