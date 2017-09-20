package main

import (
	"encoding/csv"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"path/filepath"

	"golang.org/x/text/encoding/japanese"
)

var (
	config *Config
)

func writeJson(postalCodeShort string, p areaPostal) {
	b, _ := json.Marshal(p)
	ioutil.WriteFile(filepath.Join(config.OutputDir, postalCodeShort+".json"), b, 0644)
}

func createJson() {
	fin, err := openKenAllCsv()
	if err != nil {
		panic(err)
	}
	defer fin.Close()

	jdec := japanese.ShiftJIS.NewDecoder()

	reader := csv.NewReader(jdec.Reader(fin))
	reader.LazyQuotes = true

	var phash PostalHash = make(PostalHash)

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		} else {
			p := NewPostal(record)

			if _, ok := phash[p.PostalCodeShort]; ok {
				if _, ok2 := phash[p.PostalCodeShort][p.PostalCode]; ok2 {
					phash[p.PostalCodeShort][p.PostalCode] = append(phash[p.PostalCodeShort][p.PostalCode], p)
				} else {
					phash[p.PostalCodeShort][p.PostalCode] = Postals{p}
				}
			} else {
				phash[p.PostalCodeShort] = areaPostal{
					p.PostalCode: Postals{p},
				}
			}
		}
	}

	for k, v := range phash {
		writeJson(k, v)
	}
}

func main() {
	configFileOpt := flag.String("c", "./postal.conf", "config file")
	downloadOpt := flag.Bool("d", false, "download zip")
	flag.Parse()

	var err error
	config, err = loadToml(*configFileOpt)
	if err != nil {
		fmt.Println(config)
		panic(err)
	}

	if *downloadOpt == true {
		if err := downloadPostalZip(); err != nil {
			panic(err)
		}
	}

	createJson()
}
