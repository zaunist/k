# k

<p align="center">
<a title="Go Report Card" target="_blank" href="https://goreportcard.com/report/github.com/zaunist/k"><img src="https://goreportcard.com/badge/github.com/zaunist/k" /></a>
<a title="Doc for K" target="_blank" href="https://pkg.go.dev/github.com/zaunist/k"><img src="https://pkg.go.dev/badge/github.com/zaunist/k.svg" /></a>
<a title="Release" target="_blank" href="https://github.com/zaunist/k/releases"><img src="https://img.shields.io/github/v/release/zaunist/k.svg?color=161823&style=flat-square&logo=smartthings" /></a>
<a  title="License" target="_blank" href="https://github.com/zaunist/k/blob/main/LICENSE"><img src="https://img.shields.io/github/license/zaunist/k.svg?color=161823&style=flat-square&logo=smartthings"> </a>
</p>

> 当访问不同版本的 k8s 的时候，官方的建议是 kubectl 也要跟着版本走，所以在需要操作多个跨大版本的集群时，切换 kubectl 就比较麻烦，这个工具也是因此诞生。

K 是一个用于管理 kubectl 版本的工具，可以使用它下载对应版本的 kubectl，支持使用 k use 切换至想要的版本。

## 安装

### 手动安装

```
go install github.com/zaunist/k@latest
```

### 自动化安装

```shell
curl -sSL https://mirror.ghproxy.com/https://raw.githubusercontent.com/zaunist/k/main/install.sh | bash
```

## 用法

```
K is used to manage kubectl version.

Usage:
  k [command]

Available Commands:
  clean       Remove files from the package download directory
  help        Help about any command
  install     Download and install a version
  ls          List installed versions
  uninstall   Uninstall a version of kubectl
  use         Switch to specified version
  version     Print the current version

Flags:
  -h, --help   help for k

Use "k [command] --help" for more information about a command.

```

## 示例

```shell
k install v1.23.0  // 安装 v1.23.0 版本的 kubectl

k use v1.22.0 // 切换为 v1.22.0 版本的 kubectl

k ls  // 列出当前主机上已安装 kubectl 版本
```

## 参考

[g](https://github.com/voidint/g) - Golang version manager

## 许可证

[WTFPL license](http://www.wtfpl.net/about/)
