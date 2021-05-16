package cmd

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type Context struct {
	suite.Suite
	dir string
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(Context))
}

func (ctx *Context) SetupTest() {
	ctx.dir = ctx.T().TempDir()
}
