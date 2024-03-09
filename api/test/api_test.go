package test

import (
	"testing"

	"github.com/ThCompiler/go.beget.api/core"
	"github.com/ThCompiler/ts"
	"github.com/stretchr/testify/suite"
)

type APISuite struct {
	ts.TestCasesSuite
	server *BegetServer
}

func (ap *APISuite) SetupSuite() {
	ap.server = NewBegetServer()
	core.SetMode(core.Test)
	core.SetTestHost(ap.server.GetURL())
}

func TestAPI(t *testing.T) {
	suite.Run(t, new(APISuite))
}
