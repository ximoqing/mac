# mac
[中文](README.md)|
[English](README-en.md)

[![Go Reference](https://pkg.go.dev/badge/github.com/ximoqing/mac.svg)](https://pkg.go.dev/github.com/ximoqing/mac)
[![Release](https://img.shields.io/badge/Release-0.0.1-red)](https://github.com/ximoqing/mac/releases)


#### introduce
rand mac address generation

#### install
```
go get github.com/ximoqing/mac
```

#### Example
```go
package main

import (
	"fmt"

	"github.com/ximoqing/mac"
)

func main() {
	n := mac.MacSingle()
	fmt.Println(n)
}


```