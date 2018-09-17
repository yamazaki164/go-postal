package main

import (
	"path/filepath"
	"testing"
)

func TestLoadToml(t *testing.T) {
	test1, err1 := LoadToml("/path/to/dummy")
	if err1 == nil {
		t.Fatal("illegal return error")
	}
	if test1.OutputDir != "" {
		t.Fatal("illegal OutputDir")
	}
	if test1.WorkingDir != "" {
		t.Fatal("illegal WorkingDir")
	}
	if test1.ZipUrl != "" {
		t.Fatal("illegal ZipUrl")
	}

	test2, err2 := LoadToml("./test/dummy.conf")
	if err2 != nil {
		t.Fatal("config file not found")
	}
	if test2.OutputDir != "/path/to/output/dir" {
		t.Fatal("illegal OutputDir")
	}
	if test2.WorkingDir != "/path/to/working/dir" {
		t.Fatal("illegal WorkingDir")
	}
	if test2.ZipUrl != "http://www.post.japanpost.jp/zipcode/dl/oogaki/zip/ken_all.zip" {
		t.Fatal("illegal ZipUrl")
	}
}

func TestZipName(t *testing.T) {
	test1, err := LoadToml("./test/dummy.conf")
	if err != nil {
		t.Fatal("config file not found")
	}

	if test1.ZipName() != "ken_all.zip" {
		t.Fatal("error on ZipName")
	}
}

func TestZipFile(t *testing.T) {
	test1, err := LoadToml("./test/dummy.conf")
	if err != nil {
		t.Fatal("config file not found")
	}
	if filepath.ToSlash(test1.ZipFile()) != "/path/to/working/dir/ken_all.zip" {
		t.Log(filepath.ToSlash(test1.ZipFile()))
		t.Fatal("error on ZipFile")
	}
}

func TestIsValidDir(t *testing.T) {
	conf := &Config{}
	test1 := conf.isValidDir("./test/dummydir")
	if test1 == true {
		t.Fatal("./test/dummydir exists")
	}

	test2 := conf.isValidDir("./test")
	if test2 == false {
		t.Fatal("./test not found")
	}

	test3 := conf.isValidDir("./test/dummy.conf")
	if test3 == true {
		t.Fatal("./test/dummy.conf is dir")
	}
}

func TestIsValidConfig(t *testing.T) {
	conf1 := &Config{}
	test1 := conf1.IsValidConfig()
	if test1 == true {
		t.Fatal("invalid config")
	}

	conf2 := &Config{
		WorkingDir: "./test",
		OutputDir:  "./test",
	}
	test2 := conf2.IsValidConfig()
	if test2 == false {
		t.Fatal("invalid config")
	}
}
