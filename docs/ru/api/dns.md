---
comments: true
title: "Управление DNS"
---

Функции управления DNS позволяют работать с DNS записями пользователя. Существуют два метода работы:

* CallGetData.
* CallChangeRecords.

## CallGetData

Соответствующее название метода в Beget.API: [getData](https://beget.com/ru/kb/api/funkczii-upravleniya-dns#getdata).


```go
func CallGetData(domain string) *GetDataResult
```

## CallChangeRecords

Соответствующее название метода в Beget.API: [changeRecords](https://beget.com/ru/kb/api/funkczii-upravleniya-dns#changerecords)