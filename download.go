package main

import (
	"io"
	"net/http"
	"os"
	"path/filepath"
)

const (
	zipfileUrl = "http://www.post.japanpost.jp/zipcode/dl/oogaki/zip/ken_all.zip"
	workingDir = "g:/works/tmp"
)

func zipfile() string {
	return filepath.Join(workingDir, filepath.Base(zipfileUrl))
}

func downloadPostalZip() {
	res, err := http.Get(zipfileUrl)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	file, err := os.Create(zipfile())
	if err != nil {
		panic(err)
	}
	defer file.Close()

	io.Copy(file, res.Body)
}
