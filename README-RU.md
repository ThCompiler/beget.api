[![Go Report Card](https://goreportcard.com/badge/github.com/ThCompiler/go.beget.api)](https://goreportcard.com/report/github.com/ThCompiler/go.beget.api)
[![Go Reference](https://pkg.go.dev/badge/github.com/ThCompiler/go.beget.api.svg)](https://pkg.go.dev/github.com/ThCompiler/go.beget.api)

[![en](https://img.shields.io/badge/lang-en-red.svg)](./README.md)

# Go.beget.api

Простая библиотека реализующая открытое API сервиса [beget.com](https://beget.com/) для выполнения функций панели управления.

### Установка

Для работы с библиотекой необходима версия golang: ```1.19```. Сама установка может быть произведена с помощью команды:

```cmd
go get github.com/ThCompiler/go.beget.api
```

### Поддерживаемые Beget.API функции

* Получить информацию о хостинг-аккаунте;
* Управлять резервными копиями;
* Управлять планировщиком заданий;
* Производить настройку DNS;
* Управлять базами данных;
* Создавать и удалять сайты на аккаунте;
* Управлять настройками доменов;
* Управлять почтовыми ящиками.

Детальная информация об API представлена на [сайте документации](https://beget.com/ru/kb/api/beget-api) **Beget**.

### Документация

В [документации](https://pkg.go.dev/github.com/ThCompiler/go.beget.api) представлено описание функций библиотеки. А в репозитории 
[update_hostname](https://github.com/ThCompiler/update_hostname) представлен пример использования 
библиотеки для обновления данные ip для hostname на основе текущего ip системы.

### P.S.

Вопросы и предложения можно указывать в issue данного репозитория.

Реализованное API представлено в [TODO](./TODO.md)