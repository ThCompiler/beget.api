// Package backup implements the functionality of [Backup methods].
// Package implements [getFileBackupList], [getMysqlBackupList],
// [getFileList], [getMysqlList], [restoreFile], [restoreMysql], [downloadFile], [downloadMysql], [getLog] methods.
// To create the appropriate methods, you need to call either [CallGetFileBackupList], [CallGetMysqlBackupList],
// [CallGetFileList], [CallGetMysqlList], [CallRestoreFile], [CallRestoreMysql], [CallDownloadFile],
// [CallDownloadMysql] or [CallGetLog].
//
// [Backup methods]: https://beget.com/ru/kb/api/funkczii-upravleniya-bekapami
// [getFileBackupList]: https://beget.com/ru/kb/api/funkczii-upravleniya-bekapami#getfilebackuplist
// [getMysqlBackupList]: https://beget.com/ru/kb/api/funkczii-upravleniya-bekapami#getmysqlbackuplist
// [getFileList]: https://beget.com/ru/kb/api/funkczii-upravleniya-bekapami#getfilelist
// [getMysqlList]: https://beget.com/ru/kb/api/funkczii-upravleniya-bekapami#getmysqllist
// [restoreFile]: https://beget.com/ru/kb/api/funkczii-upravleniya-bekapami#restorefile
// [restoreMysql]: https://beget.com/ru/kb/api/funkczii-upravleniya-bekapami#restoremysql
// [downloadFile]: https://beget.com/ru/kb/api/funkczii-upravleniya-bekapami#downloadfile
// [downloadMysql]: https://beget.com/ru/kb/api/funkczii-upravleniya-bekapami#downloadmysql
// [getLog]: https://beget.com/ru/kb/api/funkczii-upravleniya-bekapami#getlog
package backup

// Constants used to implement the [getFileBackupList], [getMysqlBackupList],
// [getFileList], [getMysqlList], [restoreFile], [restoreMysql], [downloadFile], [downloadMysql], [getLog] methods.
//
// [getFileBackupList]: https://beget.com/ru/kb/api/funkczii-upravleniya-bekapami#getfilebackuplist
// [getMysqlBackupList]: https://beget.com/ru/kb/api/funkczii-upravleniya-bekapami#getmysqlbackuplist
// [getFileList]: https://beget.com/ru/kb/api/funkczii-upravleniya-bekapami#getfilelist
// [getMysqlList]: https://beget.com/ru/kb/api/funkczii-upravleniya-bekapami#getmysqllist
// [restoreFile]: https://beget.com/ru/kb/api/funkczii-upravleniya-bekapami#restorefile
// [restoreMysql]: https://beget.com/ru/kb/api/funkczii-upravleniya-bekapami#restoremysql
// [downloadFile]: https://beget.com/ru/kb/api/funkczii-upravleniya-bekapami#downloadfile
// [downloadMysql]: https://beget.com/ru/kb/api/funkczii-upravleniya-bekapami#downloadmysql
// [getLog]: https://beget.com/ru/kb/api/funkczii-upravleniya-bekapami#getlog
const (
	GetFileBackupListMethodName = "GetFileBackupList"
	GetFileBackupListMethodPath = "backup/getFileBackupList"

	GetMysqlBackupListMethodName = "GetMysqlBackupList"
	GetMysqlBackupListMethodPath = "backup/getMysqlBackupList"

	GetFileListMethodName = "GetFileList"
	GetFileListMethodPath = "backup/getFileList"

	GetMysqlListMethodName = "GetMysqlList"
	GetMysqlListMethodPath = "backup/getMysqlList"

	RestoreFileMethodName = "RestoreFile"
	RestoreFileMethodPath = "backup/restoreFile"

	RestoreMysqlMethodName = "RestoreMysql"
	RestoreMysqlMethodPath = "backup/restoreMysql"

	DownloadFileMethodName = "DownloadFile"
	DownloadFileMethodPath = "backup/downloadFile"

	DownloadMysqlMethodName = "DownloadMysql"
	DownloadMysqlMethodPath = "backup/downloadMysql"

	GetLogMethodName = "GetLog"
	GetLogMethodPath = "backup/getLog"
)
