package application

import (
	mocks2 "github.com/goravel/framework/contracts/auth/mocks"
	mock2 "github.com/goravel/framework/testing/mock"
	"testing"

	"github.com/stretchr/testify/suite"
	"mgufrone.dev/job-tracking/tests"
)

type TestIndex struct {
	tests.HttpTestCase
	auth *mocks2.Auth
}

func TestIndexSuite(t *testing.T) {
	suite.Run(t, new(TestIndex))
}

// SetupTest will run before each test in the suite.
func (s *TestIndex) SetupTest() {
	s.HttpTestCase.SetupTest()
	if s.auth == nil {
		s.auth = mock2.Auth()
	}
}

// TearDownTest will run after each test in the suite.
func (s *TestIndex) TearDownTest() {
}

func (s *TestIndex) TestIndexFail0() {
}
