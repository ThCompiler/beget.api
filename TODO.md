## Реализованное API

### [Управление аккаунтом](https://beget.com/ru/kb/api/funkczii-upravleniya-akkauntom)

* [x] getAccountInfo
* [x] toggleSsh

### [Управление бэкапами](https://beget.com/ru/kb/api/funkczii-upravleniya-bekapami)

* [x] getFileBackupList
* [x] getMysqlBackupList
* [x] getFileList
* [x] getMysqlList
* [x] restoreFile
* [x] restoreMysql
* [x] downloadFile
* [x] downloadMysql
* [x] getLog

### [Управление Cron](https://beget.com/ru/kb/api/funkczii-upravleniya-cron)

* [ ] getList
* [ ] add
* [ ] edit
* [ ] delete
* [ ] changeHiddenState
* [ ] getEmail
* [ ] downloadFile
* [ ] downloadMysql
* [ ] getLog

### [Управление DNS](https://beget.com/ru/kb/api/funkczii-upravleniya-dns)

* [x] getData
* [x] changeRecords

### [Управление FTP](https://beget.com/ru/kb/api/funkczii-upravleniya-ftp)

* [ ] getList
* [ ] add
* [ ] changePassword
* [ ] delete

### [Управление MySQL](https://beget.com/ru/kb/api/funkczii-upravleniya-mysql)

* [ ] getList
* [ ] addDb
* [ ] addAccess
* [ ] dropDb
* [ ] dropAccess
* [ ] changeAccessPassword

### [Управление сайтами](https://beget.com/ru/kb/api/funkczii-upravleniya-sajtami)

* [ ] getList
* [ ] add
* [ ] delete
* [ ] linkDomain
* [ ] unlinkDomain
* [ ] freeze
* [ ] unfreeze
* [ ] isSiteFrozen
* [ ] getLog

### [Управление доменами](https://beget.com/ru/kb/api/funkczii-dlya-raboty-s-domenami)

* [ ] getList
* [ ] getZoneList
* [ ] addVirtual
* [ ] delete
* [ ] getSubdomainList
* [ ] addSubdomainVirtual
* [ ] deleteSubdomain
* [ ] checkDomainToRegister
* [ ] getPhpVersion
* [ ] changePhpVersion
* [ ] getDirectives
* [ ] addDirectives
* [ ] removeDirectives

### [Управление почтой](https://beget.com/ru/kb/api/funkczii-dlya-raboty-s-pochtoj)

* [ ] getMailboxList
* [ ] changeMailboxPassword
* [ ] createMailbox
* [ ] dropMailbox
* [ ] changeMailboxSettings
* [ ] forwardListAddMailbox
* [ ] forwardListDeleteMailbox
* [ ] forwardListShow
* [ ] setDomainMail
* [ ] clearDomainMail

### [Сбор статистики](https://beget.com/ru/kb/api/funkczii-dlya-sbora-statistiki)

* [ ] getSiteListLoad
* [ ] getDbListLoad
* [ ] getSiteLoad
* [ ] getDbLoad
