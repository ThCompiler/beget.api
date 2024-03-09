package result

import (
	"strings"

	"github.com/pkg/errors"

	"github.com/ThCompiler/go.beget.api/internal/time"
)

var ErrUnknownValueForFileType = errors.New("unknown value fot type FileType") // "unknown value fot type FileType"

// FileBackup is an entry in the list of file backups.
type FileBackup struct {
	BackupID ID             `json:"backup_id"`
	Date     time.BegetTime `json:"date"`
}

// MYSQLBackup is an entry in the list of MYSQL backups.
type MYSQLBackup FileBackup

// FileBackupList is the result of a successful call to the [getFileBackupList] method.
//
// [getFileBackupList]: https://beget.com/ru/kb/api/funkczii-upravleniya-bekapami#getfilebackuplist
type FileBackupList []FileBackup

// MYSQLBackupList is the result of a successful call to the [getMysqlBackupList] method.
//
// [getMysqlBackupList]: https://beget.com/ru/kb/api/funkczii-upravleniya-bekapami#getmysqlbackuplist
type MYSQLBackupList []MYSQLBackup

// FileType represents the type of file on the ftp server. Means true if it is a directory, otherwise it is a file.
type FileType bool

// UnmarshalJSON specifies the unmarshalling of FileType from a JSON string.
// FileType is true if the Json string contains "true" or "1".
// FileType is false if the Json string contains "false", "null", "1" or the string is empty.
func (ft *FileType) UnmarshalJSON(data []byte) error {
	switch strings.Trim(string(data), "\" ") {
	case "true", "1":
		*ft = true
	case "false", "0", "null", "":
		*ft = false
	default:
		return errors.Wrapf(ErrUnknownValueForFileType, "with data %s", data)
	}

	return nil
}

// FileRecord is an entry in the list of files in a directory.
type FileRecord struct {
	Name         string         `json:"name"`
	IsDirectory  FileType       `json:"is_dir"`
	CreationTime time.BegetTime `json:"mtime"`
	Size         uint64         `json:"size"`
}

// FileList is the result of a successful call to the [getFileList] method.
//
// [getFileList]: https://beget.com/ru/kb/api/funkczii-upravleniya-bekapami#getfilelist
type FileList []FileRecord

// MYSQLList is the result of a successful call to the [getMysqlList] method.
//
// [getMysqlList]: https://beget.com/ru/kb/api/funkczii-upravleniya-bekapami#getmysqllist
type MYSQLList []DatabaseName

// Operation represents the operation performed on the database or files in the Beget.
type Operation string

const (
	Restore  Operation = "restore"  // files or databases war restored.
	Download Operation = "download" // files or databases war downloaded.
)

// OperationType represents a specific operation performed in the Beget.
type OperationType string

const (
	RestoreFile   OperationType = "restore_file"   // files war restored.
	DownloadFile  OperationType = "download_file"  // files war downloaded.
	RestoreMYSQL  OperationType = "restore_mysql"  // databases war restored.
	DownloadMYSQL OperationType = "download_mysql" // databases war downloaded.
)

// Status represents the status of the operation performed in the Beget.
type Status string

const (
	SUCCESS Status = "success" // the operation was completed successfully.
	ERROR   Status = "error"   // the operation failed with an error.
)

// LogRecord is a record in the operation log in the Beget.
type LogRecord struct {
	ID         ID             `json:"id"`
	Operation  Operation      `json:"operation"`
	Type       OperationType  `json:"type"`
	DateCreate time.BegetTime `json:"date_create"`
	TargetList []string       `json:"target_list"`
	Status     Status         `json:"status"`
}

// Log is the result of a successful call to the [getLog] method.
//
// [getLog]: https://beget.com/ru/kb/api/funkczii-upravleniya-bekapami#getlog
type Log []LogRecord
