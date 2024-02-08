---
comments: true
title: "Установка"
---

Библиотека разработана под версию golang: ```1.19```. 

#### Скачивание

Библиотека может быть установлена единственным способ: через ```go get```:

```cmd
go get github.com/ThCompiler/github.com/ThCompiler/go.beget.api
```

#### Использования

После чего можно использовать её в своём проекта подключив в качестве пакета:

``` go linenums="1" hl_lines="7-10
package main

import (
	"fmt"
	"log"

	"github.com/ThCompiler/go.beget.api/pkg/beget/api/dns"
	"github.com/ThCompiler/go.beget.api/pkg/beget/api/dns/build"
	"github.com/ThCompiler/go.beget.api/pkg/beget/api/result"
	"github.com/ThCompiler/go.beget.api/pkg/beget/core"
)
```

Полный код представлен в разделе [Пример](./example.md)