# K - Kubectl version manager

[English](README.md) | [简体中文](README.zh.md)

## Overview

K is used to manage kubectl version, ff you need to control k8s across large versions(v1.18 and v1.23), then this is the right tool for you.

## Getting Started

```shell
go install github.com/zaunist/k@latest
```

Download in [release page](https://github.com/zaunist/k/releases)

## Usage

```
k is used to manage kubectl version.

Usage:
  k [command]

Available Commands:
  clean       Clean all resource you download
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  install     install special kubectl version with --version
  ls          ls local kubectl version
  uninstall   Uninstall kubectl version
  use         use special version kubectl
  version     The K version

Flags:
  -h, --help   help for k

Use "k [command] --help" for more information about a command.

```

k install v1.23.0  install v1.23.0 version

k use v1.22.0  switch kubectl version to v1.22.0

k ls  list all installed version on your system

## Reference

This project inspire by [g](https://github.com/voidint/g) - Golang version manager

## License

[WTFPL license](LICENSE)
