package main

import (
	"encoding/csv"
	"encoding/json"
	"io"
	"io/ioutil"
	"path/filepath"

	"github.com/yamazaki164/go-postal/postal"

	"golang.org/x/text/encoding/japanese"
)

func writeJson(postalCodeShort string, p postal.AreaPostal) {
	b, _ := json.Marshal(p)
	ioutil.WriteFile(filepath.Join(config.OutputDir, postalCodeShort+".json"), b, 0644)
}

func createJson(phash postal.PostalHash) {
	for k, v := range phash {
		writeJson(k, v)
	}
}

func Parse() postal.PostalHash {
	fin, err := openKenAllCsv()
	if err != nil {
		panic(err)
	}
	defer fin.Close()

	jdec := japanese.ShiftJIS.NewDecoder()

	reader := csv.NewReader(jdec.Reader(fin))
	reader.LazyQuotes = true

	var phash postal.PostalHash = make(postal.PostalHash)

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		} else {
			p := postal.NewPostal(record)

			if _, ok := phash[p.PostalCodeShort]; ok {
				if _, ok2 := phash[p.PostalCodeShort][p.PostalCode]; ok2 {
					phash[p.PostalCodeShort][p.PostalCode] = append(phash[p.PostalCodeShort][p.PostalCode], p)
				} else {
					phash[p.PostalCodeShort][p.PostalCode] = postal.Postals{p}
				}
			} else {
				phash[p.PostalCodeShort] = postal.AreaPostal{
					p.PostalCode: postal.Postals{p},
				}
			}
		}
	}

	return phash
}
