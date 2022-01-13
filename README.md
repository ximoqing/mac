# mac
[中文](README.md)|
[English](README-en.md)

![GoDoc](https://pkg.go.dev/badge/github.com/ximoqing/mac?status.svg)
[![Release](https://img.shields.io/badge/Release-0.0.1-red)](https://github.com/ximoqing/mac/releases)

#### 介绍
mac地址生成

#### 安装
```
go get github.com/ximoqing/mac
```

#### 示例
```go
func main() {
    mac := mac.MacInit()
    fmt.Println(mac)
}

```

