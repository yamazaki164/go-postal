package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"

	"strconv"

	"golang.org/x/text/encoding/japanese"
)

type Postal struct {
	JisCode         string `json:"jis_code"`
	OldPostalCode   string `json:"old_postal_code"`
	PostalCode      string `json:"postal_code"`
	PostalCodeShort string `json:"postal_code_short"`
	KanaPrefecture  string `json:"kana_prefecture"`
	KanaAddress1    string `json:"kana_address1"`
	KanaAddress2    string `json:"kana_address2"`
	Prefecture      string `json:"prefecture"`
	Address1        string `json:"address1"`
	Address2        string `json:"address2"`
	Flag1           bool   `json:"flag1"`
	Flag2           bool   `json:"flag2"`
	Flag3           bool   `json:"flag3"`
	Flag4           bool   `json:"flag4"`
	Status          int64  `json:"status"`
	Reason          int64  `json:"reason"`
}

func NewPostal(r []string) *Postal {
	flag1, _ := strconv.ParseBool(r[9])
	flag2, _ := strconv.ParseBool(r[10])
	flag3, _ := strconv.ParseBool(r[11])
	flag4, _ := strconv.ParseBool(r[12])
	status, _ := strconv.ParseInt(r[13], 10, 32)
	reason, _ := strconv.ParseInt(r[14], 10, 32)

	p := &Postal{
		JisCode:        r[0],
		OldPostalCode:  r[1],
		PostalCode:     r[2],
		KanaPrefecture: r[3],
		KanaAddress1:   r[4],
		KanaAddress2:   r[5],
		Prefecture:     r[6],
		Address1:       r[7],
		Address2:       r[8],
		Flag1:          flag1,
		Flag2:          flag2,
		Flag3:          flag3,
		Flag4:          flag4,
		Status:         status,
		Reason:         reason,
	}

	return p
}

func main() {
	fin, err := os.OpenFile("g:/tmp/KEN_ALL.CSV", os.O_RDONLY, 0755)
	if err != nil {
		panic(err.Error())
	}
	defer fin.Close()

	jdec := japanese.ShiftJIS.NewDecoder()

	cin := csv.NewReader(jdec.Reader(fin))
	cin.LazyQuotes = true
	postalList := []*Postal{}
	for {
		record, err := cin.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		} else {
			p := NewPostal(record)
			postalList = append(postalList, p)
		}
	}

	fmt.Println(len(postalList))
}
