package api

import "github.com/stretchr/testify/require"

func (ctx *Context) TestCreateNamespace() {
	assert := require.New(ctx.T())

	err := ctx.api.CreateNamespace(
		"testproject",
		"testgroup",
		"test description",
	)
	assert.Nil(err)

	expectedPaths := []string{
		"cluster-scope/base/core/namespaces/testproject/kustomization.yaml",
		"cluster-scope/base/core/namespaces/testproject/namespace.yaml",
	}

	compareWithExpected(assert, "testdata/CreateNamespace", ctx.dir, expectedPaths)
}
