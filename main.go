package main

import (
	"encoding/csv"
	"encoding/json"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"

	"golang.org/x/text/encoding/japanese"
)

func writeJson(k1 string, p areaPostal) {
	var rootDir = filepath.Join("c:/works/tmp/", k1)

	_, err := os.Stat(rootDir)
	if err != nil && os.IsNotExist(err) {
		os.MkdirAll(rootDir, 0644)
	}

	b, _ := json.Marshal(p)
	ioutil.WriteFile(filepath.Join(rootDir, k1+".json"), b, 0644)
}

func createJson() {
	fin, err := os.OpenFile(zipfile(), os.O_RDONLY, 0755)
	if err != nil {
		panic(err.Error())
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
	//downloadPostalZip()
	createJson()
}
