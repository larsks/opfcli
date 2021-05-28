package api

import "github.com/stretchr/testify/require"

func (ctx *Context) TestCreateRoleBinding() {
	assert := require.New(ctx.T())

	err := ctx.api.CreateRoleBinding(
		"testgroup",
		"admin",
	)
	assert.Nil(err)

	expectedPaths := []string{
		"cluster-scope/components/project-admin-rolebindings/testgroup/rbac.yaml",
		"cluster-scope/components/project-admin-rolebindings/testgroup/kustomization.yaml",
	}

	compareWithExpected(assert, "testdata/CreateRoleBinding", ctx.dir, expectedPaths)
}
