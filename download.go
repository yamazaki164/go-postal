package main

import (
	"io"
	"net/http"
	"os"
)

func downloadPostalZip() error {
	res, err := http.Get(config.ZipUrl)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	file, err := os.Create(config.Zipfile())
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, res.Body)
	return err
}
