package main

import (
	"errors"
	"flag"
	"fmt"
)

var (
	config        *Config
	configFileOpt *string
	downloadOpt   *bool
	silentOpt     *bool
)

func ParseFlags() {
	configFileOpt = flag.String("c", "./postal.conf", "/path/to/config/file")
	downloadOpt = flag.Bool("download", false, "download zip. (default: not download)")
	silentOpt = flag.Bool("s", false, "silent mode")
	flag.Parse()
}

func main() {
	ParseFlags()
	var err error
	config, err = LoadToml(*configFileOpt)
	if err != nil {
		panic(err)
	}

	if !config.IsValidConfig() {
		panic(errors.New("invalid config settings"))
	}

	if *downloadOpt == true {
		if !*silentOpt {
			fmt.Println("#download " + config.ZipName())
		}
		if err := downloadPostalZip(); err != nil {
			panic(err)
		}
	}

	if !*silentOpt {
		fmt.Println("#parse and generate postal json")
	}

	createJson(Parse())
}
