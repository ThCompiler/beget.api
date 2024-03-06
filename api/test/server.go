package test

import (
	"github.com/ThCompiler/go.beget.api/api/dns"
	"github.com/ThCompiler/go.beget.api/api/user"
	"github.com/ThCompiler/go.beget.api/core"
	"github.com/ThCompiler/go.beget.api/pkg/maps"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

type HandlerPair struct {
	Body                string
	CheckValuesFunction func(url.Values)
}

func Default() *HandlerPair {
	return &HandlerPair{
		Body:                "",
		CheckValuesFunction: func(url.Values) {},
	}
}

func APIHandler(pair *HandlerPair) http.HandlerFunc {
	return func(resp http.ResponseWriter, req *http.Request) {
		pair.CheckValuesFunction(req.URL.Query())

		_, _ = resp.Write([]byte(pair.Body))

		resp.WriteHeader(http.StatusOK)
	}
}

type Handlers struct {
	GetAccountInfo *HandlerPair
	ToggleSSH      *HandlerPair
	GetData        *HandlerPair
	ChangeRecords  *HandlerPair
}

func DefaultHandlers() *Handlers {
	return &Handlers{
		GetAccountInfo: Default(),
		ToggleSSH:      Default(),
		GetData:        Default(),
		ChangeRecords:  Default(),
	}
}

type BegetServer struct {
	handlers *Handlers
	mux      *http.ServeMux
	server   *httptest.Server
}

func NewBegetServer() *BegetServer {
	handlers := DefaultHandlers()
	mux := http.NewServeMux()

	mux.HandleFunc("/"+user.GetAccountInfoPath, APIHandler(handlers.GetAccountInfo))
	mux.HandleFunc("/"+user.ToggleSSHMethodPath, APIHandler(handlers.ToggleSSH))
	mux.HandleFunc("/"+dns.GetDataMethodPath, APIHandler(handlers.GetData))
	mux.HandleFunc("/"+dns.ChangeRecordsMethodPath, APIHandler(handlers.ChangeRecords))

	return &BegetServer{
		handlers: handlers,
		mux:      mux,
		server:   httptest.NewServer(mux),
	}
}

func (s *BegetServer) UserGetAccountInfo(body string, checkValuesFunction func(url.Values)) *BegetServer {
	s.handlers.GetAccountInfo.Body = body
	s.handlers.GetAccountInfo.CheckValuesFunction = checkValuesFunction
	return s
}

func (s *BegetServer) UserToggleSSH(body string, checkValuesFunction func(url.Values)) *BegetServer {
	s.handlers.ToggleSSH.Body = body
	s.handlers.ToggleSSH.CheckValuesFunction = checkValuesFunction
	return s
}

func (s *BegetServer) DNSGetData(body string, checkValuesFunction func(url.Values)) *BegetServer {
	s.handlers.GetData.Body = body
	s.handlers.GetData.CheckValuesFunction = checkValuesFunction
	return s
}

func (s *BegetServer) DNSChangeRecords(body string, checkValuesFunction func(url.Values)) *BegetServer {
	s.handlers.ChangeRecords.Body = body
	s.handlers.ChangeRecords.CheckValuesFunction = checkValuesFunction
	return s
}

func (s *BegetServer) GetURL() string {
	return s.server.URL
}

func (s *BegetServer) Start() {
	s.server.Start()
}

func (s *BegetServer) Stop() {
	s.server.Close()
}

func RequireEqualValues(t *testing.T, expected, actual url.Values, client core.Client, skipValues ...string) {
	expected.Add("login", client.Login)
	expected.Add("passwd", client.Password)
	expected.Add("output_format", string(core.JSON))

	cpExpected := maps.Clone(expected)
	cpActual := maps.Clone(actual)

	for _, values := range skipValues {
		cpExpected.Del(values)
		cpActual.Del(values)
	}

	require.EqualValues(t, cpExpected, cpActual)
}
