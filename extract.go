package main

import (
	"archive/zip"
	"errors"
	"io"
)

const (
	targetCsvName = "KEN_ALL.CSV"
)

func openKenAllCsv() (io.ReadCloser, error) {
	zreader, err := zip.OpenReader(config.ZipFile())
	if err != nil {
		return nil, errors.New("zip open error: " + err.Error())
	}

	for _, f := range zreader.File {
		if f.FileInfo().IsDir() {
			continue
		}

		if f.FileInfo().Name() == targetCsvName {
			return f.Open()
		}
	}

	return nil, errors.New("CSV not found")
}
