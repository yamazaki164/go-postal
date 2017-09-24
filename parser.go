package main

import (
	"encoding/csv"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"path/filepath"

	"github.com/yamazaki164/go-postal/postal"

	"golang.org/x/text/encoding/japanese"
)

func writeJson(postalCodeShort string, p postal.AreaPostal) error {
	if p == nil {
		return errors.New("nil pointer error")
	}

	b, e := json.Marshal(p)
	if e != nil {
		return e
	}

	return ioutil.WriteFile(filepath.Join(config.OutputDir, postalCodeShort+".json"), b, 0644)
}

func createJson(phash postal.PostalHash) {
	if phash == nil {
		panic(errors.New("nil hash error"))
	}

	for k, v := range phash {
		if e := writeJson(k, v); e != nil {
			panic(e)
		}
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
