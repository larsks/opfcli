package api

import "github.com/stretchr/testify/require"

func (ctx *Context) TestCreateGroup() {
	assert := require.New(ctx.T())

	err := ctx.api.CreateGroup(
		"testgroup",
		false,
	)
	assert.Nil(err)

	expectedPaths := []string{
		"cluster-scope/base/user.openshift.io/groups/testgroup/group.yaml",
		"cluster-scope/base/user.openshift.io/groups/testgroup/kustomization.yaml",
	}

	compareWithExpected(assert, "testdata/CreateGroup", ctx.dir, expectedPaths)

	// ---

	// Should fail if group already exists and existsOk is false
	err = ctx.api.CreateGroup(
		"testgroup",
		false,
	)
	assert.EqualError(err, "group testgroup already exists")

	// ---

	// Should work if group already exists and existsOk is true
	err = ctx.api.CreateGroup(
		"testgroup",
		true,
	)
	assert.Nil(err)
}
