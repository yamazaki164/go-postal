package postal

import (
	"strconv"
	"strings"
)

type Postal struct {
	JisCode         string `json:"jis_code"`
	OldPostalCode   string `json:"-"`
	PostalCode      string `json:"postal_code"`
	PostalCodeShort string `json:"-"`
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
	Status          int64  `json:"-"`
	Reason          int64  `json:"-"`
}

func NewPostal(r []string) *Postal {
	flag1, _ := strconv.ParseBool(strings.TrimSpace(r[9]))
	flag2, _ := strconv.ParseBool(strings.TrimSpace(r[10]))
	flag3, _ := strconv.ParseBool(strings.TrimSpace(r[11]))
	flag4, _ := strconv.ParseBool(strings.TrimSpace(r[12]))
	status, _ := strconv.ParseInt(strings.TrimSpace(r[13]), 10, 32)
	reason, _ := strconv.ParseInt(strings.TrimSpace(r[14]), 10, 32)

	p := &Postal{
		JisCode:         strings.TrimSpace(r[0]),
		OldPostalCode:   strings.TrimSpace(r[1]),
		PostalCode:      strings.TrimSpace(r[2]),
		PostalCodeShort: strings.TrimSpace(r[2])[0:3],
		KanaPrefecture:  strings.TrimSpace(r[3]),
		KanaAddress1:    strings.TrimSpace(r[4]),
		KanaAddress2:    strings.TrimSpace(r[5]),
		Prefecture:      strings.TrimSpace(r[6]),
		Address1:        strings.TrimSpace(r[7]),
		Address2:        strings.TrimSpace(r[8]),
		Flag1:           flag1,
		Flag2:           flag2,
		Flag3:           flag3,
		Flag4:           flag4,
		Status:          status,
		Reason:          reason,
	}

	return p
}

type Postals []*Postal

type AreaPostal map[string]Postals

type PostalHash map[string]AreaPostal
