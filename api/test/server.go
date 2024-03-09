// Package test implements the functionality of test server emulating the work of a Beget server.
package test

import (
	"net/http"
	"net/http/httptest"
	"net/url"

	"github.com/ThCompiler/go.beget.api/api/backup"
	"github.com/ThCompiler/go.beget.api/api/dns"
	"github.com/ThCompiler/go.beget.api/api/user"
	"github.com/pkg/errors"
)

var ErrNotSetHandler = errors.New("handler wasn't set") // "handler wasn't set".

type handlerPair struct {
	Body                string
	CheckValuesFunction func(url.Values)
}

func emptyPair() *handlerPair {
	return &handlerPair{
		Body:                "",
		CheckValuesFunction: nil,
	}
}

func (h *handlerPair) Set(body string, checkValuesFunction func(url.Values)) {
	h.Body = body
	h.CheckValuesFunction = checkValuesFunction
}

func aPIHandler(pair *handlerPair) http.HandlerFunc {
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

type handlers struct {
	GetAccountInfo     *handlerPair
	ToggleSSH          *handlerPair
	GetData            *handlerPair
	ChangeRecords      *handlerPair
	GetFileBackupList  *handlerPair
	GetFileList        *handlerPair
	GetMYSQLBackupList *handlerPair
	GetMYSQLList       *handlerPair
	RestoreFile        *handlerPair
	RestoreMYSQL       *handlerPair
	DownloadFile       *handlerPair
	DownloadMYSQL      *handlerPair
	GetLog             *handlerPair
}

func defaultHandlers() *handlers {
	return &handlers{
		GetAccountInfo:     emptyPair(),
		ToggleSSH:          emptyPair(),
		GetData:            emptyPair(),
		ChangeRecords:      emptyPair(),
		GetFileBackupList:  emptyPair(),
		GetFileList:        emptyPair(),
		GetMYSQLBackupList: emptyPair(),
		GetMYSQLList:       emptyPair(),
		RestoreFile:        emptyPair(),
		RestoreMYSQL:       emptyPair(),
		DownloadFile:       emptyPair(),
		DownloadMYSQL:      emptyPair(),
		GetLog:             emptyPair(),
	}
}

// BegetServer is an extension of the [httptest.Server] to work with Beget Server.
// Supports processing of all api methods implemented by the package.
//
// # Important
//
// If the handler is not specified and the appropriate method is invoked,
// the server will panic with [ErrNotSetHandler].
type BegetServer struct {
	handlers *handlers
	mux      *http.ServeMux
	server   *httptest.Server
}

// NewBegetServer creates a new [BegetServer].
// The [httptest.Server] will be started.
func NewBegetServer() *BegetServer {
	handlers := defaultHandlers()
	mux := http.NewServeMux()

	mux.HandleFunc("/"+user.GetAccountInfoPath, aPIHandler(handlers.GetAccountInfo))
	mux.HandleFunc("/"+user.ToggleSSHMethodPath, aPIHandler(handlers.ToggleSSH))
	mux.HandleFunc("/"+dns.GetDataMethodPath, aPIHandler(handlers.GetData))
	mux.HandleFunc("/"+dns.ChangeRecordsMethodPath, aPIHandler(handlers.ChangeRecords))
	mux.HandleFunc("/"+backup.GetFileBackupListMethodPath, aPIHandler(handlers.GetFileBackupList))
	mux.HandleFunc("/"+backup.GetMysqlBackupListMethodPath, aPIHandler(handlers.GetMYSQLBackupList))
	mux.HandleFunc("/"+backup.GetFileListMethodPath, aPIHandler(handlers.GetFileList))
	mux.HandleFunc("/"+backup.GetMysqlListMethodPath, aPIHandler(handlers.GetMYSQLList))
	mux.HandleFunc("/"+backup.GetLogMethodPath, aPIHandler(handlers.GetLog))
	mux.HandleFunc("/"+backup.RestoreFileMethodPath, aPIHandler(handlers.RestoreFile))
	mux.HandleFunc("/"+backup.RestoreMysqlMethodPath, aPIHandler(handlers.RestoreMYSQL))
	mux.HandleFunc("/"+backup.DownloadFileMethodPath, aPIHandler(handlers.DownloadFile))
	mux.HandleFunc("/"+backup.DownloadMysqlMethodPath, aPIHandler(handlers.DownloadMYSQL))

	return &BegetServer{
		handlers: handlers,
		mux:      mux,
		server:   httptest.NewServer(mux),
	}
}

// DNSGetData specify handler for [getData] method.
//
// [getData]: https://beget.com/ru/kb/api/funkczii-upravleniya-dns#getdata
func (s *BegetServer) DNSGetData(body string, checkValuesFunction func(url.Values)) *BegetServer {
	s.handlers.GetData.Set(body, checkValuesFunction)

	return s
}

// DNSChangeRecords specify handler for [changeRecords] method.
//
// [changeRecords]: https://beget.com/ru/kb/api/funkczii-upravleniya-dns#changerecords
func (s *BegetServer) DNSChangeRecords(body string, checkValuesFunction func(url.Values)) *BegetServer {
	s.handlers.ChangeRecords.Set(body, checkValuesFunction)

	return s
}

// UserGetAccountInfo specify handler for [getAccountInfo] method.
//
// [getAccountInfo]: https://beget.com/ru/kb/api/funkczii-upravleniya-akkauntom#getaccountinfo
func (s *BegetServer) UserGetAccountInfo(body string, checkValuesFunction func(url.Values)) *BegetServer {
	s.handlers.GetAccountInfo.Set(body, checkValuesFunction)

	return s
}

// UserToggleSSH specify handler for [toggleSsh] method.
//
// [toggleSsh]: https://beget.com/ru/kb/api/funkczii-upravleniya-akkauntom#togglessh
func (s *BegetServer) UserToggleSSH(body string, checkValuesFunction func(url.Values)) *BegetServer {
	s.handlers.ToggleSSH.Set(body, checkValuesFunction)

	return s
}

// BackupGetFileBackupList specify handler for [getFileBackupList] method.
//
// [getFileBackupList]: https://beget.com/ru/kb/api/funkczii-upravleniya-bekapami#getfilebackuplist
func (s *BegetServer) BackupGetFileBackupList(body string, checkValuesFunction func(url.Values)) *BegetServer {
	s.handlers.GetFileBackupList.Set(body, checkValuesFunction)

	return s
}

// BackupGetMYSQLBackupList specify handler for [getMysqlBackupList] method.
//
// [getMysqlBackupList]: https://beget.com/ru/kb/api/funkczii-upravleniya-bekapami#getmysqlbackuplist
func (s *BegetServer) BackupGetMYSQLBackupList(body string, checkValuesFunction func(url.Values)) *BegetServer {
	s.handlers.GetMYSQLBackupList.Set(body, checkValuesFunction)

	return s
}

// BackupGetFileList specify handler for [getFileList] method.
//
// [getFileList]: https://beget.com/ru/kb/api/funkczii-upravleniya-bekapami#getfilelist
func (s *BegetServer) BackupGetFileList(body string, checkValuesFunction func(url.Values)) *BegetServer {
	s.handlers.GetFileList.Set(body, checkValuesFunction)

	return s
}

// BackupGetMYSQLList specify handler for [getMysqlList] method.
//
// [getMysqlList]: https://beget.com/ru/kb/api/funkczii-upravleniya-bekapami#getmysqllist
func (s *BegetServer) BackupGetMYSQLList(body string, checkValuesFunction func(url.Values)) *BegetServer {
	s.handlers.GetMYSQLList.Set(body, checkValuesFunction)

	return s
}

// BackupGetLog specify handler for [getLog] method.
//
// [getLog]: https://beget.com/ru/kb/api/funkczii-upravleniya-bekapami#getlog
func (s *BegetServer) BackupGetLog(body string, checkValuesFunction func(url.Values)) *BegetServer {
	s.handlers.GetLog.Set(body, checkValuesFunction)

	return s
}

// BackupRestoreFile specify handler for [restoreFile] method.
//
// [restoreFile]: https://beget.com/ru/kb/api/funkczii-upravleniya-bekapami#restorefile
func (s *BegetServer) BackupRestoreFile(body string, checkValuesFunction func(url.Values)) *BegetServer {
	s.handlers.RestoreFile.Set(body, checkValuesFunction)

	return s
}

// BackupRestoreMYSQL specify handler for [restoreMysql] method.
//
// [restoreMysql]: https://beget.com/ru/kb/api/funkczii-upravleniya-bekapami#restoremysql
func (s *BegetServer) BackupRestoreMYSQL(body string, checkValuesFunction func(url.Values)) *BegetServer {
	s.handlers.RestoreMYSQL.Set(body, checkValuesFunction)

	return s
}

// BackupDownloadFile specify handler for [downloadFile] method.
//
// [downloadFile]: https://beget.com/ru/kb/api/funkczii-upravleniya-bekapami#downloadfile
func (s *BegetServer) BackupDownloadFile(body string, checkValuesFunction func(url.Values)) *BegetServer {
	s.handlers.DownloadFile.Set(body, checkValuesFunction)

	return s
}

// BackupDownloadMYSQL specify handler for [downloadMysql] method.
//
// [downloadMysql]: https://beget.com/ru/kb/api/funkczii-upravleniya-bekapami#downloadmysql
func (s *BegetServer) BackupDownloadMYSQL(body string, checkValuesFunction func(url.Values)) *BegetServer {
	s.handlers.DownloadMYSQL.Set(body, checkValuesFunction)

	return s
}

// GetURL returns URL of test server.
func (s *BegetServer) GetURL() string {
	return s.server.URL
}
