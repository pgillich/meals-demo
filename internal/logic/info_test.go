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
	"github.com/pgillich/meals-demo/internal/api"
)

type InfoTestSuite struct {
	suite.Suite
}

func TestInfoTestSuite(t *testing.T) {
	os.Args = os.Args[:1]
	suite.Run(t, new(InfoTestSuite))
}

func (s *InfoTestSuite) buildServerInfo(options *configs.Options) *http.Server {
	return internal.BuildServer(options, NewFoodStore)
}

func (s *InfoTestSuite) TestGetLivez() {
	os.Setenv("PORT", "8081")
	options := configs.Options{
		DbDialect: "sqlite3",
		DbDsn:     ":memory:",
		DbSample:  true,
		DbDebug:   true,
		JwtKey:    "1234",
	}
	server := s.buildServerInfo(&options)
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", configs.TestingBasePath+"/livez", nil)

	server.Handler.ServeHTTP(w, r)
	defer server.Close()

	s.Equal(http.StatusOK, w.Result().StatusCode, "Status")
	s.Equal("", w.Body.String(), "Body")
}

func (s *InfoTestSuite) TestGetVersion() {
	os.Setenv("PORT", "8081")
	options := configs.Options{
		DbDialect: "sqlite3",
		DbDsn:     ":memory:",
		DbSample:  true,
		DbDebug:   true,
		JwtKey:    "1234",
	}
	server := s.buildServerInfo(&options)
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", configs.TestingBasePath+"/version", nil)

	server.Handler.ServeHTTP(w, r)
	defer server.Close()

	s.Equal(http.StatusOK, w.Result().StatusCode, "Status")
	version := api.Version{}
	err := json.Unmarshal(w.Body.Bytes(), &version)
	s.NoError(err, "Body")
}
