package logic

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/pgillich/meals-demo/configs"
	"github.com/pgillich/meals-demo/internal"
	"github.com/pgillich/meals-demo/internal/models"
	"github.com/pgillich/meals-demo/internal/restapi"
)

type InfoTestSuite struct {
	suite.Suite
}

func TestInfoTestSuite(t *testing.T) {
	os.Args = os.Args[:1]
	suite.Run(t, new(InfoTestSuite))
}

func (s *InfoTestSuite) buildServerInfo(options *configs.Options) *restapi.Server {
	return internal.BuildServer(options, SetInfoAPI)
}

func (s *InfoTestSuite) TestGetLivez() {
	options := configs.Options{}
	server := s.buildServerInfo(&options)
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", configs.TestingBasePath+"/livez", nil)

	server.GetHandler().ServeHTTP(w, r)

	s.Equal(http.StatusOK, w.Result().StatusCode, "Status")
	s.Equal("", w.Body.String(), "Body")
}

func (s *InfoTestSuite) TestGetVersion() {
	options := configs.Options{}
	server := s.buildServerInfo(&options)
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", configs.TestingBasePath+"/version", nil)

	server.GetHandler().ServeHTTP(w, r)

	s.Equal(http.StatusOK, w.Result().StatusCode, "Status")
	version := models.Version{}
	err := json.Unmarshal(w.Body.Bytes(), &version)
	s.NoError(err, "Body")
}
