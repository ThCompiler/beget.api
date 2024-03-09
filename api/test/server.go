package test

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/ThCompiler/go.beget.api/api/backup"
	"github.com/ThCompiler/go.beget.api/api/dns"
	"github.com/ThCompiler/go.beget.api/api/user"
	"github.com/ThCompiler/go.beget.api/core"
	"github.com/ThCompiler/go.beget.api/pkg/maps"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"
)

var ErrNotSetHandler = errors.New("handler wasn't set")

type HandlerPair struct {
	Body                string
	CheckValuesFunction func(url.Values)
}

func Default() *HandlerPair {
	return &HandlerPair{
		Body:                "",
		CheckValuesFunction: nil,
	}
}

func (h *HandlerPair) Set(body string, checkValuesFunction func(url.Values)) {
	h.Body = body
	h.CheckValuesFunction = checkValuesFunction
}

func APIHandler(pair *HandlerPair) http.HandlerFunc {
	return func(resp http.ResponseWriter, req *http.Request) {
		if pair.CheckValuesFunction == nil {
			panic(errors.Wrapf(ErrNotSetHandler, "with path of handler %s", req.URL.Path))
		}

		pair.CheckValuesFunction(req.URL.Query())

		_, err := resp.Write([]byte(pair.Body))
		if err != nil {
			panic(errors.Wrapf(err, "when write body %s with path of handler %s", pair.Body, req.URL.Path))
		}
	}
}

type Handlers struct {
	GetAccountInfo     *HandlerPair
	ToggleSSH          *HandlerPair
	GetData            *HandlerPair
	ChangeRecords      *HandlerPair
	GetFileBackupList  *HandlerPair
	GetFileList        *HandlerPair
	GetMYSQLBackupList *HandlerPair
	GetMYSQLList       *HandlerPair
	RestoreFile        *HandlerPair
	RestoreMYSQL       *HandlerPair
	DownloadFile       *HandlerPair
	DownloadMYSQL      *HandlerPair
	GetLog             *HandlerPair
}

func DefaultHandlers() *Handlers {
	return &Handlers{
		GetAccountInfo:     Default(),
		ToggleSSH:          Default(),
		GetData:            Default(),
		ChangeRecords:      Default(),
		GetFileBackupList:  Default(),
		GetFileList:        Default(),
		GetMYSQLBackupList: Default(),
		GetMYSQLList:       Default(),
		RestoreFile:        Default(),
		RestoreMYSQL:       Default(),
		DownloadFile:       Default(),
		DownloadMYSQL:      Default(),
		GetLog:             Default(),
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
	mux.HandleFunc("/"+backup.GetFileBackupListMethodPath, APIHandler(handlers.GetFileBackupList))
	mux.HandleFunc("/"+backup.GetMysqlBackupListMethodPath, APIHandler(handlers.GetMYSQLBackupList))
	mux.HandleFunc("/"+backup.GetFileListMethodPath, APIHandler(handlers.GetFileList))
	mux.HandleFunc("/"+backup.GetMysqlListMethodPath, APIHandler(handlers.GetMYSQLList))
	mux.HandleFunc("/"+backup.GetLogMethodPath, APIHandler(handlers.GetLog))
	mux.HandleFunc("/"+backup.RestoreFileMethodPath, APIHandler(handlers.RestoreFile))
	mux.HandleFunc("/"+backup.RestoreMysqlMethodPath, APIHandler(handlers.RestoreMYSQL))
	mux.HandleFunc("/"+backup.DownloadFileMethodPath, APIHandler(handlers.DownloadFile))
	mux.HandleFunc("/"+backup.DownloadMysqlMethodPath, APIHandler(handlers.DownloadMYSQL))

	return &BegetServer{
		handlers: handlers,
		mux:      mux,
		server:   httptest.NewServer(mux),
	}
}

func (s *BegetServer) UserGetAccountInfo(body string, checkValuesFunction func(url.Values)) *BegetServer {
	s.handlers.GetAccountInfo.Set(body, checkValuesFunction)

	return s
}

func (s *BegetServer) UserToggleSSH(body string, checkValuesFunction func(url.Values)) *BegetServer {
	s.handlers.ToggleSSH.Set(body, checkValuesFunction)

	return s
}

func (s *BegetServer) DNSGetData(body string, checkValuesFunction func(url.Values)) *BegetServer {
	s.handlers.GetData.Set(body, checkValuesFunction)

	return s
}

func (s *BegetServer) DNSChangeRecords(body string, checkValuesFunction func(url.Values)) *BegetServer {
	s.handlers.ChangeRecords.Set(body, checkValuesFunction)

	return s
}

func (s *BegetServer) BackupGetFileBackupList(body string, checkValuesFunction func(url.Values)) *BegetServer {
	s.handlers.GetFileBackupList.Set(body, checkValuesFunction)

	return s
}

func (s *BegetServer) BackupGetMYSQLBackupList(body string, checkValuesFunction func(url.Values)) *BegetServer {
	s.handlers.GetMYSQLBackupList.Set(body, checkValuesFunction)

	return s
}

func (s *BegetServer) BackupGetFileList(body string, checkValuesFunction func(url.Values)) *BegetServer {
	s.handlers.GetFileList.Set(body, checkValuesFunction)

	return s
}

func (s *BegetServer) BackupGetMYSQLList(body string, checkValuesFunction func(url.Values)) *BegetServer {
	s.handlers.GetMYSQLList.Set(body, checkValuesFunction)

	return s
}

func (s *BegetServer) BackupGetLog(body string, checkValuesFunction func(url.Values)) *BegetServer {
	s.handlers.GetLog.Set(body, checkValuesFunction)

	return s
}

func (s *BegetServer) BackupRestoreFile(body string, checkValuesFunction func(url.Values)) *BegetServer {
	s.handlers.RestoreFile.Set(body, checkValuesFunction)

	return s
}

func (s *BegetServer) BackupRestoreMYSQL(body string, checkValuesFunction func(url.Values)) *BegetServer {
	s.handlers.RestoreMYSQL.Set(body, checkValuesFunction)

	return s
}

func (s *BegetServer) BackupDownloadFile(body string, checkValuesFunction func(url.Values)) *BegetServer {
	s.handlers.DownloadFile.Set(body, checkValuesFunction)

	return s
}

func (s *BegetServer) BackupDownloadMYSQL(body string, checkValuesFunction func(url.Values)) *BegetServer {
	s.handlers.DownloadMYSQL.Set(body, checkValuesFunction)

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
	t.Helper()

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
