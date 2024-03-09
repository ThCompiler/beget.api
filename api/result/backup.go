package result

import (
	"github.com/pkg/errors"

	"github.com/ThCompiler/go.beget.api/internal/time"
)

var ErrUnknownValueForFileType = errors.New("unknown value fot type FileType")

type ID uint64

type FileBackup struct {
	BackupID ID             `json:"backup_id"`
	Date     time.BegetTime `json:"date"`
}

type MYSQLBackup FileBackup

type FileBackupList []FileBackup

type MYSQLBackupList []MYSQLBackup

type FileType bool

func (ft *FileType) UnmarshalJSON(data []byte) error {
	switch string(data) {
	case "true", "1":
		*ft = true
	case "false", "0", "null", "":
		*ft = false
	default:
		return errors.Wrapf(ErrUnknownValueForFileType, "with data %s", data)
	}

	return nil
}

type FileRecord struct {
	Name         string         `json:"name"`
	IsDirectory  FileType       `json:"is_dir"`
	CreationTime time.BegetTime `json:"mtime"`
	Size         uint64         `json:"size"`
}

type FileList []FileRecord

type DatabaseName string

type MYSQLList []DatabaseName

type Operation string

const (
	Restore  Operation = "restore"
	Download Operation = "download"
)

type OperationType string

const (
	RestoreFile   OperationType = "restore_file"
	DownloadFile  OperationType = "download_file"
	RestoreMYSQL  OperationType = "restore_mysql"
	DownloadMYSQL OperationType = "download_mysql"
)

type Status string

const (
	SUCCESS Status = "success"
	ERROR   Status = "error"
)

type LogRecord struct {
	ID         ID             `json:"id"`
	Operation  Operation      `json:"operation"`
	Type       OperationType  `json:"type"`
	DateCreate time.BegetTime `json:"date_create"`
	TargetList []string       `json:"target_list"`
	Status     Status         `json:"status"`
}

type Log []LogRecord
