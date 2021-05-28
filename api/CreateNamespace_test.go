package api

import "github.com/stretchr/testify/require"

func (ctx *Context) TestCreateNamespace() {
	assert := require.New(ctx.T())

	err := ctx.api.CreateNamespace(
		"testproject",
		"testgroup",
		"test description",
		"",
		false,
	)
	assert.Nil(err)

	expectedPaths := []string{
		"cluster-scope/base/core/namespaces/testproject/kustomization.yaml",
		"cluster-scope/base/core/namespaces/testproject/namespace.yaml",
	}

	compareWithExpected(assert, "testdata/CreateNamespace", ctx.dir, expectedPaths)
}

func (ctx *Context) TestCreateNamespaceQuota() {
	assert := require.New(ctx.T())

	// Should fail if quota doesn't exist
	err := ctx.api.CreateNamespace(
		"testproject",
		"testgroup",
		"test description",
		"testquota",
		false,
	)
	assert.Nil(err)

	expectedPaths := []string{
		"cluster-scope/base/core/namespaces/testproject/kustomization.yaml",
	}

	compareWithExpected(assert, "testdata/CreateNamespaceQuota", ctx.dir, expectedPaths)
}

func (ctx *Context) TestCreateNamespaceNoLimitrange() {
	assert := require.New(ctx.T())

	// Should fail if quota doesn't exist
	err := ctx.api.CreateNamespace(
		"testproject",
		"testgroup",
		"test description",
		"",
		true,
	)
	assert.Nil(err)

	expectedPaths := []string{
		"cluster-scope/base/core/namespaces/testproject/kustomization.yaml",
	}

	compareWithExpected(assert, "testdata/CreateNamespaceNoLimitrange", ctx.dir, expectedPaths)
}
