package main

import (
	"regexp"
	"testing"

	"github.com/yamazaki164/go-postal/postal"
)

func TestWriteJson(t *testing.T) {
	config = &Config{
		WorkingDir: "./test",
		OutputDir:  "./test",
	}

	test1 := writeJson("003", nil)
	if a, _ := regexp.MatchString("nil pointer error", test1.Error()); !a {
		t.Log(test1)
		t.Fatal("not null pointer on writeJson")
	}

	p := &postal.Postal{}
	pl := postal.Postals{p}
	data2 := postal.AreaPostal{
		"0010002": pl,
	}

	test2 := writeJson("003", data2)
	if test2 != nil {
		t.Log(test1)
		t.Fatal(test1)
	}
}

func TestCreateJson(t *testing.T) {
	func() {
		defer func() {
			err := recover()
			if err == nil {
				t.Fatal("nil hash error")
			}
		}()

		createJson(nil)
	}()

	func() {
		defer func() {
			err := recover()
			if err != nil {
				t.Fatal("error")
			}
		}()

		p := &postal.Postal{}
		ph := postal.PostalHash{
			"001": postal.AreaPostal{
				"0010002": postal.Postals{p, p},
			},
		}
		createJson(ph)
	}()
}

func TestParse(t *testing.T) {
	config = &Config{
		WorkingDir: "./test",
		OutputDir:  "./test",
		ZipUrl:     "testdata3.zip",
	}

	test1 := Parse()
	_, b := test1["060"]
	if !b {
		t.Log(test1)
		t.Fatal("parse error on correct file")
	}

	func() {
		defer func() {
			err := recover()
			if err == nil {
				t.Fatal("parse error on incorrect file")
			}
		}()

		config = &Config{
			WorkingDir: "./test",
			OutputDir:  "./test",
		}
		test2 := Parse()
		if test2 == nil {
			t.Fatal("parse error on incorrect file")
		}
	}()
}
