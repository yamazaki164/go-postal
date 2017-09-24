package main

import (
	"testing"
)

func TestParseFlags(t *testing.T) {
	configFileOpt = nil
	downloadOpt = nil
	silentOpt = nil

	ParseFlags()
	if *configFileOpt != "./postal.conf" {
		t.Fatal("configFileOpt parse error")
	}

	if *downloadOpt {
		t.Fatal("downloadOpt parse error")
	}

	if *silentOpt {
		t.Fatal("silentOpt parse error")
	}
}

func TestMain(t *testing.T) {
	t.Skip()
}
