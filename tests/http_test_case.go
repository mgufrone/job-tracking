package tests

import (
	"github.com/goravel/framework/facades"
	"github.com/goravel/framework/testing"
	"github.com/stretchr/testify/suite"
	"net/http/httptest"
)

type HttpTestCase struct {
	suite.Suite
	testing.TestCase
	ApiServer *httptest.Server
}

func (h *HttpTestCase) SetupTest() {
	if h.ApiServer == nil {
		h.ApiServer = httptest.NewServer(facades.Route())
	}
}
