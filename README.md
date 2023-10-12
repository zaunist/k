# K - Kubectl Version Manager

<p align="center">
<a title="Go Report Card" target="_blank" href="https://goreportcard.com/report/github.com/zaunist/k"><img src="https://goreportcard.com/badge/github.com/zaunist/k" /></a>
<a title="Doc for K" target="_blank" href="https://pkg.go.dev/github.com/zaunist/k"><img src="https://pkg.go.dev/badge/github.com/zaunist/k.svg" /></a>
<a title="Release" target="_blank" href="https://github.com/zaunist/k/releases"><img src="https://img.shields.io/github/v/release/zaunist/k.svg?color=161823&style=flat-square&logo=smartthings" /></a>
<a  title="License" target="_blank" href="https://github.com/zaunist/k/blob/main/LICENSE"><img src="https://img.shields.io/github/license/zaunist/k.svg?color=161823&style=flat-square&logo=smartthings"> </a>
</p>

[English](README.md) | [简体中文](README.zh.md)

## Overview

K is used to manage kubectl version, if you need to control k8s across large versions(v1.18 and v1.23), then this is the right tool for you.

## Getting Started

### manual

```shell
go install github.com/zaunist/k@latest
```

### auto

```shell
curl -sSL https://raw.githubusercontent.com/zaunist/k/main/install.sh | bash
```

## Usage

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

## example

```
k install v1.23.0  // install v1.23.0 version

k use v1.22.0 //  switch kubectl version to v1.22.0

k ls // list all installed version on your system

```

## Reference

The project is inspired by [g](https://github.com/voidint/g) - Golang version manager

## License

[WTFPL license](http://www.wtfpl.net/about/)
