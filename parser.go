package main

import (
	"encoding/csv"
	"encoding/json"
	"io"
	"io/ioutil"
	"path/filepath"

	"golang.org/x/text/encoding/japanese"
)

func writeJson(postalCodeShort string, p areaPostal) {
	b, _ := json.Marshal(p)
	ioutil.WriteFile(filepath.Join(config.OutputDir, postalCodeShort+".json"), b, 0644)
}

func createJson(phash PostalHash) {
	for k, v := range phash {
		writeJson(k, v)
	}
}

func Parse() PostalHash {
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

	return phash
}
