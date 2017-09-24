package main

import (
	"testing"
)

func TestDownloadPostalZip(t *testing.T) {
	config = &Config{
		WorkingDir: "./test",
		ZipUrl:     "http://www.post.japanpost.jp/zipcode/dl/oogaki/zip/ken_all.zip",
	}
	test1 := downloadPostalZip()
	if test1 != nil {
		t.Log(config)
		t.Fatal("error on download")
	}

	config = &Config{}
	test2 := downloadPostalZip()
	if test2 == nil {
		t.Fatal("error on failure download")
	}
}
