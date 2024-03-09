## 0.0.1

## Добавлено

### Библиотека

#### Ядро

* Добавлены типы ошибок, форматы и статусы (раздел core/info).
* Добавлена структура описание ответа от сервисом api.
* Добавлено обобщённое выполнения запросов по api и получение результатов.
* Добавлен интерфейс `ApiMethod` описывающий требования к методом api.

#### Api

* Добавлен базовый метод описывающий запрос по api.
* Добавлены методы взаимодействия с DNS: `getData`, `changeRecords`.
* Реализован создатель записей для запроса на изменение dns записей через метод `changeRecords`.

### Репозиторий

* Добавлено описание README на русском и английском языке.
* Настроена работа CI.
* Добавлен пример применения библиотеки в другом репозитории.

## 0.0.2

## Исправлено

### Библиотека

#### Ядро

* Исправлено возвращение значения ***nil*** для ошибки из метода `Get` у `BegetResponse` и `Answer`.

### Репозиторий

* Исправлен путь к документации в README.
* Исправлено путь получения пакета через команду `go get` в README.

## Изменено

### Библиотека

#### Ядро

* Переименованы методы `GetResult` и `MustGetResult` типа `BegetResponse` на `GetAnswer` и `MustGetAnswer`, соответственно.
* Переработана структура пакетов.

#### Api

* Сделан экспортируемым интерфейс `settableRecords` в пакете **api/dns**.
* Изменены тип поля **records** в методе `CallChangeRecords` на `SettableRecords` в пакете **api/dns**.
* Удалены ненужные функции `SetRecords`, `SetBasicRecords`, `SetNsRecords`, `SetCNameRecords` в пакете **api/dns**.
* Удалены ненужный тип `SettingRecords` в пакете **api/dns**.
* Переименован метод `NewARecordCreator` из пакета **api/dns/build** на `NewARecords`.
* Переименован метод `NewAAAARecordCreator` из пакета **api/dns/build** на `NewAAAARecords`.
* Переименован метод `NewMxRecordCreator` из пакета **api/dns/build** на `NewMxRecords`.
* Переименован метод `NewTxtRecordCreator` из пакета **api/dns/build** на `NewTxtRecords`.
* Переименован метод `NewNsRecordCreator` из пакета **api/dns/build** на `NewNsRecords`.
* Переименован метод `NewCNameRecordCreator` из пакета **api/dns/build** на `NewCNameRecords`.
* Переименован метод `NewDNSIPRecordCreator` из пакета **api/dns/build** на `NewDNSIPRecords`.
* Переименован метод `NewDNSRecordCreator` из пакета **api/dns/build** на `NewDNSRecords`.
* Изменено встраивание типа `DNSRecordsCreator` в `BasicRecordsCreator`, `NsRecordsCreator`, `CNameRecordsCreator`
  на хранение поля **dnsRecords** с типом `DNSRecordsCreator` в пакетe **api/dns/build**.

## Добавлено

### Библиотека

#### Ядро

* Добавлена документация Golang.

#### Api

* Добавлена документация Golang.

## 0.1.1

## Изменено

### Библиотека

#### Api

* Добавлено получение `url.Values` в качестве параметра метода `CallMethod` структуры `BasicMethod` в пакете **api**. 
  И, соответственно, появилась возможность добавлять в информацию о методе API параметры запроса.

## Добавлено

### Библиотека

#### Ядро

* Добавлен метод `PrepareRequestWithClient` в пакете **core**, позволяющий указывать пользовательский 
  `http.Client` для запросов к API.
* Добавлено тестовое состояние для работы API.

#### Api

* Добавлен пакет тест в пакет *api* с тестами API методов.
* Добавлены методы [управления аккаунтами](https://beget.com/ru/kb/api/funkczii-upravleniya-akkauntom) в 
  пакет **api/user**.
* Добавлены методы [управления бэкапами](https://beget.com/ru/kb/api/funkczii-upravleniya-bekapami) в
  пакет **api/backup**.

### Pkg

* Добавлена функция для клонирования `map`.