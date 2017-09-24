package postal

import (
	"testing"
)

func TestNewPostal(t *testing.T) {
	test1 := []string{
		"01101",
		"060  ",
		"0600000",
		"ﾎﾂｶｲﾄﾞｳ",
		"ｻﾂﾎﾟﾛｼﾁﾕｳｵｳｸ",
		"ｲｶﾆｹｲｻｲｶﾞﾅｲﾊﾞｱｲ",
		"北海道",
		"札幌市中央区",
		"以下に掲載がない場合",
		"0",
		"0",
		"0",
		"0",
		"0",
		"0",
	}

	result1 := NewPostal(test1)
	if result1.JisCode != "01101" {
		t.Fatal("parse error on jis")
	}
	if result1.OldPostalCode != "060" {
		t.Fatal("parse error on old_postal_code")
	}
	if result1.PostalCode != "0600000" {
		t.Fatal("parse error on postal_code")
	}
	if result1.PostalCodeShort != "060" {
		t.Fatal("parse error on postal_code_short")
	}
	if result1.KanaPrefecture != "ﾎﾂｶｲﾄﾞｳ" {
		t.Fatal("parse error on kana prefecture")
	}
	if result1.KanaAddress1 != "ｻﾂﾎﾟﾛｼﾁﾕｳｵｳｸ" {
		t.Fatal("parse error on kana address1")
	}
	if result1.KanaAddress2 != "ｲｶﾆｹｲｻｲｶﾞﾅｲﾊﾞｱｲ" {
		t.Fatal("parse error on kana address2")
	}
	if result1.Prefecture != "北海道" {
		t.Fatal("parse error on prefecture")
	}
	if result1.Address1 != "札幌市中央区" {
		t.Fatal("parse error on address1")
	}
	if result1.Address2 != "以下に掲載がない場合" {
		t.Fatal("parse error on address2")
	}
	if result1.Flag1 == true {
		t.Fatal("parse error on flag1")
	}
	if result1.Flag2 == true {
		t.Fatal("parse error on flag2")
	}
	if result1.Flag3 == true {
		t.Fatal("parse error on flag3")
	}
	if result1.Flag4 == true {
		t.Fatal("parse error on flag4")
	}
	if result1.Reason != 0 {
		t.Fatal("parse error on reason")
	}
	if result1.Status != 0 {
		t.Fatal("parse error on status")
	}
}
