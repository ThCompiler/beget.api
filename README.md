[![Go Report Card](https://goreportcard.com/badge/github.com/ThCompiler/go.beget.api)](https://goreportcard.com/report/github.com/ThCompiler/go.beget.api)
[![Go Reference](https://pkg.go.dev/badge/github.com/ThCompiler/go.beget.api.svg)](https://pkg.go.dev/github.com/ThCompiler/go.beget.api)


[![en](https://img.shields.io/badge/lang-ru-green.svg)](./README-RU.md)

# Go.beget.api

A simple library that implements the open API of the service [baget.com](https://beget.com/)
to perform the functions of the control panel.

### Install

To work with the library, you need the golang version: ``1.19``. The installation can be done using the command:

```cmd
go get github.com/ThCompiler/go.beget.api
```

### Functions supported by the Beget.API

* Get information about your hosting account;
* Manage backups;
* Manage the task scheduler;
* Configure DNS;
* Manage databases;
* Create and delete websites on your account;
* Manage domain settings;
* Manage mailboxes.

Detailed information about the API is provided on the **Beget`s** [documentation site](https://beget.com/ru/kb/api/beget-api).

### Documentation

The [documentation]https://pkg.go.dev/github.com/ThCompiler/go.beget.api) provides a description of the library's functions.
And the [update_hostname](https://github.com/ThCompiler/update_hostname) repository provides an example of usage libraries
for updating ip for hostname based on the current ip of the server.

### P.S.

Questions and suggestions can be specified in the issue of this repository.