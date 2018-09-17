package main

import (
	"regexp"
	"testing"
)

func TestOpenKenAllCsv(t *testing.T) {
	config = &Config{}
	_, err1 := openKenAllCsv()
	if a, _ := regexp.MatchString("zip open error: ", err1.Error()); !a {
		t.Log(err1)
		t.Fatal("error on open ZipFile")
	}

	config = &Config{
		WorkingDir: "./test",
		OutputDir:  "./test",
		ZipUrl:     "testdata.zip",
	}
	_, err2 := openKenAllCsv()
	if a, _ := regexp.MatchString("CSV not found", err2.Error()); !a {
		t.Log(err2)
		t.Fatal("read error on OpenReader")
	}

	config = &Config{
		WorkingDir: "./test",
		OutputDir:  "./test",
		ZipUrl:     "testdata2.zip",
	}
	_, err3 := openKenAllCsv()
	if a, _ := regexp.MatchString("CSV not found", err3.Error()); !a {
		t.Log(err3)
		t.Fatal("read error on OpenReader")
	}

	config = &Config{
		WorkingDir: "./test",
		OutputDir:  "./test",
		ZipUrl:     "ken_all.zip",
	}
	test4, err4 := openKenAllCsv()
	if err4 != nil {
		t.Log(err4)
		t.Fatal("error at open inner file")
	}
	if test4 == nil {
		t.Fatal("error at open inner file")
	}
}
