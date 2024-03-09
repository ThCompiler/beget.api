package test

import (
	"encoding/json"
	"net/url"

	"github.com/ThCompiler/go.beget.api/api/backup"
	"github.com/ThCompiler/go.beget.api/api/result"
	"github.com/ThCompiler/go.beget.api/core"
	"github.com/ThCompiler/go.beget.api/internal/time"
	"github.com/stretchr/testify/require"
)

var backupFileListResponse = result.FileBackupList{
	{BackupID: 390059403, Date: time.MustParse("2024-03-06 07:49:41")},
	{BackupID: 389902697, Date: time.MustParse("2024-03-04 08:02:47")},
	{BackupID: 389691984, Date: time.MustParse("2024-03-01 14:40:05")},
}

const getBackupListResponse = `
{
    "status": "success",
    "answer": {
        "status": "success",
        "result": [
			{{ range $index, $file := . }}{{ if $index }},{{end}}
			{
				"backup_id": {{ $file.BackupID }},
                "date": "{{ $file.Date }}"
			}
			{{ end }}
        ]
    }
}
`

func (ap *APISuite) TestGetBackupFileList() {
	client := core.Client{
		Login:    "login",
		Password: "password",
	}

	ap.server.BackupGetFileBackupList(
		createByTemplate(ap.T(), "fileBackup", getBackupListResponse, backupFileListResponse),
		func(values url.Values) {
			requireEqualValues(ap.T(), url.Values{}, values, client)
		},
	)

	req, err := core.PrepareRequest[result.FileBackupList](
		client,
		backup.CallGetFileBackupList(),
	)
	require.NoError(ap.T(), err)

	resp, err := req.Do()
	require.NoError(ap.T(), err)

	answ, err := resp.Get()
	require.NoError(ap.T(), err)

	res, err := answ.Get()
	require.NoError(ap.T(), err)

	require.EqualValues(ap.T(), backupFileListResponse, *res)
}

var backupMYSQLListResponse = result.MYSQLBackupList{
	{BackupID: 390059403, Date: time.MustParse("2024-03-06 07:49:41")},
	{BackupID: 389902697, Date: time.MustParse("2024-03-04 08:02:47")},
	{BackupID: 389691984, Date: time.MustParse("2024-03-01 14:40:05")},
}

func (ap *APISuite) TestGetBackupMYSQLList() {
	client := core.Client{
		Login:    "login",
		Password: "password",
	}

	ap.server.BackupGetMYSQLBackupList(
		createByTemplate(ap.T(), "mysqlBackup", getBackupListResponse, backupMYSQLListResponse),
		func(values url.Values) {
			requireEqualValues(ap.T(), url.Values{}, values, client)
		},
	)

	req, err := core.PrepareRequest[result.MYSQLBackupList](
		client,
		backup.CallGetMysqlBackupList(),
	)
	require.NoError(ap.T(), err)

	resp, err := req.Do()
	require.NoError(ap.T(), err)

	answ, err := resp.Get()
	require.NoError(ap.T(), err)

	res, err := answ.Get()
	require.NoError(ap.T(), err)

	require.EqualValues(ap.T(), backupMYSQLListResponse, *res)
}

var fileListResponse = result.FileList{
	{Name: "cgi-bin", IsDirectory: true, CreationTime: time.MustParse("2024-03-06 07:49:41"), Size: 4096},
	{Name: "index.php", IsDirectory: false, CreationTime: time.MustParse("2024-03-06 08:49:41"), Size: 5982},
}

const getFileListResponse = `
{
    "status": "success",
    "answer": {
        "status": "success",
        "result": [
			{{ range $index, $file := . }}{{ if $index }},{{end}}
			 {
                "name": "{{ $file.Name }}",
                "is_dir": {{ $file.IsDirectory }},
                "mtime": "{{ $file.CreationTime }}",
                "size": {{ $file.Size }}
            }
			{{ end }}
        ]
    }
}
`

type fileListRequest struct {
	BackupID result.ID `json:"backup_id,omitempty"`
	Path     string    `json:"path"`
}

func (ap *APISuite) TestGetFileList() {
	backupID := result.ID(665667354)
	path := "/site/public_html"
	client := core.Client{
		Login:    "login",
		Password: "password",
	}

	ap.server.BackupGetFileList(
		createByTemplate(ap.T(), "files", getFileListResponse, fileListResponse),
		func(values url.Values) {
			requireEqualValues(ap.T(), url.Values{
				"input_format": []string{string(core.JSON)},
			}, values, client, "input_data")
			require.True(ap.T(), values.Has("input_data"))
			request := values.Get("input_data")

			var req fileListRequest
			require.NoError(ap.T(), json.Unmarshal([]byte(request), &req))

			require.EqualValues(ap.T(), fileListRequest{BackupID: backupID, Path: path}, req)
		},
	)

	req, err := core.PrepareRequest[result.FileList](
		client,
		backup.CallGetFileList(&backupID, path),
	)
	require.NoError(ap.T(), err)

	resp, err := req.Do()
	require.NoError(ap.T(), err)

	answ, err := resp.Get()
	require.NoError(ap.T(), err)

	res, err := answ.Get()
	require.NoError(ap.T(), err)

	require.EqualValues(ap.T(), fileListResponse, *res)
}

var mYSQLListResponse = result.MYSQLList{
	"db1",
	"db2",
}

const getMYSQLListResponse = `
{
    "status": "success",
    "answer": {
        "status": "success",
        "result": [
			{{ range $index, $dbName := . }}{{ if $index }},{{end}}
            "{{ $dbName }}"
			{{ end }}
        ]
    }
}
`

type mYSQLListRequest struct {
	BackupID result.ID `json:"backup_id,omitempty"`
}

func (ap *APISuite) TestGetMYSQLList() {
	backupID := result.ID(665667354)
	client := core.Client{
		Login:    "login",
		Password: "password",
	}

	ap.server.BackupGetMYSQLList(
		createByTemplate(ap.T(), "MYSQLes", getMYSQLListResponse, mYSQLListResponse),
		func(values url.Values) {
			requireEqualValues(ap.T(), url.Values{
				"input_format": []string{string(core.JSON)},
			}, values, client, "input_data")
			require.True(ap.T(), values.Has("input_data"))
			request := values.Get("input_data")

			var req mYSQLListRequest
			require.NoError(ap.T(), json.Unmarshal([]byte(request), &req))

			require.EqualValues(ap.T(), mYSQLListRequest{BackupID: backupID}, req)
		},
	)

	req, err := core.PrepareRequest[result.MYSQLList](
		client,
		backup.CallGetMysqlList(&backupID),
	)
	require.NoError(ap.T(), err)

	resp, err := req.Do()
	require.NoError(ap.T(), err)

	answ, err := resp.Get()
	require.NoError(ap.T(), err)

	res, err := answ.Get()
	require.NoError(ap.T(), err)

	require.EqualValues(ap.T(), mYSQLListResponse, *res)
}

const backupSuccessResponse = `
{
    "status": "success",
    "answer": {
        "status": "success",
        "result": true
    }
}
`

type restoreFileRequest struct {
	BackupID result.ID `json:"backup_id"`
	Paths    []string  `json:"paths"`
}

func (ap *APISuite) TestRestoreFile() {
	backupID := result.ID(665667354)
	paths := []string{"/site/public_html", "/site2/public_html"}

	client := core.Client{
		Login:    "login",
		Password: "password",
	}

	ap.server.BackupRestoreFile(
		backupSuccessResponse,
		func(values url.Values) {
			requireEqualValues(ap.T(), url.Values{
				"input_format": []string{string(core.JSON)},
			}, values, client, "input_data")
			require.True(ap.T(), values.Has("input_data"))
			request := values.Get("input_data")

			var req restoreFileRequest
			require.NoError(ap.T(), json.Unmarshal([]byte(request), &req))

			require.EqualValues(ap.T(), restoreFileRequest{BackupID: backupID, Paths: paths}, req)
		},
	)

	req, err := core.PrepareRequest[result.BoolResult](
		client,
		backup.CallRestoreFile(backupID, paths),
	)
	require.NoError(ap.T(), err)

	resp, err := req.Do()
	require.NoError(ap.T(), err)

	answ, err := resp.Get()
	require.NoError(ap.T(), err)

	res, err := answ.Get()
	require.NoError(ap.T(), err)

	require.True(ap.T(), bool(*res))
}

type restoreMYSQLRequest struct {
	BackupID result.ID             `json:"backup_id"`
	Bases    []result.DatabaseName `json:"bases"`
}

func (ap *APISuite) TestRestoreMYSQL() {
	backupID := result.ID(665667354)
	bases := []result.DatabaseName{"db1", "db2"}

	client := core.Client{
		Login:    "login",
		Password: "password",
	}

	ap.server.BackupRestoreMYSQL(
		backupSuccessResponse,
		func(values url.Values) {
			requireEqualValues(ap.T(), url.Values{
				"input_format": []string{string(core.JSON)},
			}, values, client, "input_data")
			require.True(ap.T(), values.Has("input_data"))
			request := values.Get("input_data")

			var req restoreMYSQLRequest
			require.NoError(ap.T(), json.Unmarshal([]byte(request), &req))

			require.EqualValues(ap.T(), restoreMYSQLRequest{BackupID: backupID, Bases: bases}, req)
		},
	)

	req, err := core.PrepareRequest[result.BoolResult](
		client,
		backup.CallRestoreMysql(backupID, bases),
	)
	require.NoError(ap.T(), err)

	resp, err := req.Do()
	require.NoError(ap.T(), err)

	answ, err := resp.Get()
	require.NoError(ap.T(), err)

	res, err := answ.Get()
	require.NoError(ap.T(), err)

	require.True(ap.T(), bool(*res))
}

type downloadFileRequest struct {
	BackupID result.ID `json:"backup_id,omitempty"`
	Paths    []string  `json:"paths"`
}

func (ap *APISuite) TestDownloadFile() {
	backupID := result.ID(665667354)
	paths := []string{"/site/public_html", "/site2/public_html"}

	client := core.Client{
		Login:    "login",
		Password: "password",
	}

	ap.server.BackupDownloadFile(
		backupSuccessResponse,
		func(values url.Values) {
			requireEqualValues(ap.T(), url.Values{
				"input_format": []string{string(core.JSON)},
			}, values, client, "input_data")
			require.True(ap.T(), values.Has("input_data"))
			request := values.Get("input_data")

			var req downloadFileRequest
			require.NoError(ap.T(), json.Unmarshal([]byte(request), &req))

			require.EqualValues(ap.T(), downloadFileRequest{BackupID: backupID, Paths: paths}, req)
		},
	)

	req, err := core.PrepareRequest[result.BoolResult](
		client,
		backup.CallDownloadFile(&backupID, paths),
	)
	require.NoError(ap.T(), err)

	resp, err := req.Do()
	require.NoError(ap.T(), err)

	answ, err := resp.Get()
	require.NoError(ap.T(), err)

	res, err := answ.Get()
	require.NoError(ap.T(), err)

	require.True(ap.T(), bool(*res))
}

type downloadMYSQLRequest struct {
	BackupID result.ID             `json:"backup_id,omitempty"`
	Bases    []result.DatabaseName `json:"bases"`
}

func (ap *APISuite) TestDownloadMYSQL() {
	backupID := result.ID(665667354)
	bases := []result.DatabaseName{"db1", "db2"}

	client := core.Client{
		Login:    "login",
		Password: "password",
	}

	ap.server.BackupDownloadMYSQL(
		backupSuccessResponse,
		func(values url.Values) {
			requireEqualValues(ap.T(), url.Values{
				"input_format": []string{string(core.JSON)},
			}, values, client, "input_data")
			require.True(ap.T(), values.Has("input_data"))
			request := values.Get("input_data")

			var req downloadMYSQLRequest
			require.NoError(ap.T(), json.Unmarshal([]byte(request), &req))

			require.EqualValues(ap.T(), downloadMYSQLRequest{BackupID: backupID, Bases: bases}, req)
		},
	)

	req, err := core.PrepareRequest[result.BoolResult](
		client,
		backup.CallDownloadMysql(&backupID, bases),
	)
	require.NoError(ap.T(), err)

	resp, err := req.Do()
	require.NoError(ap.T(), err)

	answ, err := resp.Get()
	require.NoError(ap.T(), err)

	res, err := answ.Get()
	require.NoError(ap.T(), err)

	require.True(ap.T(), bool(*res))
}

var logResponse = result.Log{
	{
		ID: 2746366, Operation: result.Download, Type: result.DownloadMYSQL,
		DateCreate: time.MustParse("2024-03-06 07:49:41"), TargetList: []string{"bd1", "bd2"}, Status: result.SUCCESS,
	},
	{
		ID: 2778466, Operation: result.Restore, Type: result.RestoreMYSQL,
		DateCreate: time.MustParse("2024-03-06 12:49:41"), TargetList: []string{"bd5", "bd4"}, Status: result.SUCCESS,
	},
	{
		ID: 2746126, Operation: result.Download, Type: result.DownloadFile,
		DateCreate: time.MustParse("2024-03-01 07:49:41"),
		TargetList: []string{"/site/public_html", "/site2/public_html"}, Status: result.SUCCESS,
	},
	{
		ID: 2746378, Operation: result.Restore, Type: result.RestoreFile,
		DateCreate: time.MustParse("2024-03-05 07:49:41"),
		TargetList: []string{"/site/public_html", "/site2/public_html"}, Status: result.ERROR,
	},
}

const getLogResponse = `
{
    "status": "success",
    "answer": {
        "status": "success",
        "result": [
			{{ range $index, $record := . }}{{ if $index }},{{end}}
			{
				"id": {{ $record.ID }},
                "operation": "{{ $record.Operation }}",
                "type": "{{ $record.Type }}",
                "date_create": "{{ $record.DateCreate }}",
                "target_list": [
					{{ range $indexNest, $target := $record.TargetList }}{{ if $indexNest }},{{end}}
                    "{{ $target }}"
					{{ end }}
                ],
                "status": "{{ $record.Status }}"
            }
			{{ end }}
        ]
    }
}
`

func (ap *APISuite) TestGetLog() {
	client := core.Client{
		Login:    "login",
		Password: "password",
	}

	ap.server.BackupGetLog(
		createByTemplate(ap.T(), "log", getLogResponse, logResponse),
		func(values url.Values) {
			requireEqualValues(ap.T(), url.Values{}, values, client, "input_data")
		},
	)

	req, err := core.PrepareRequest[result.Log](
		client,
		backup.CallGetLog(),
	)
	require.NoError(ap.T(), err)

	resp, err := req.Do()
	require.NoError(ap.T(), err)

	answ, err := resp.Get()
	require.NoError(ap.T(), err)

	res, err := answ.Get()
	require.NoError(ap.T(), err)

	require.EqualValues(ap.T(), logResponse, *res)
}
