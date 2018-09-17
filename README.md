go-postal
====

[![CircleCI](https://circleci.com/gh/yamazaki164/go-postal/tree/master.svg?style=svg)](https://circleci.com/gh/yamazaki164/go-postal/tree/master)
[![codecov](https://codecov.io/gh/yamazaki164/go-postal/branch/master/graph/badge.svg)](https://codecov.io/gh/yamazaki164/go-postal)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)


Overview

郵便番号データ加工プログラム

## Description

日本郵便のWebサイトで提供されている郵便番号データをダウンロード・パースし、郵便番号の先頭3桁ごとにJSONファイルとして出力する。

## Requirement

- Go: 1.9+
- golang.org/x/text/encoding/japanese
- github.com/BurntSushi/toml

## Installation

```
go get github.com/yamazaki164/go-postal
```

## Usage

```shell
$ ./go-postal -c /path/to/config -download
```

```shell
$ ./go-postal -h
Usage of go-postal:
  -c string
        /path/to/config/file (default "./postal.conf")

  -download
        download zip. (default: not download)
  -s    silent mode
```

## Configration

|param|value type|description|
|:--|:--|:--|
| output_dir | string | output directory of json files |
| working_dir | string | download directory of jp-postal-zip file |
| zip_url | string | URL of jp-postal-zip |


## Licence

MIT