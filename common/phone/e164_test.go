package phone

import (
	"github.com/ttacon/libphonenumber"
	"testing"
)

func TestPhone1(t *testing.T) {

	p1, err := libphonenumber.Parse("+8613312341234", "")
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(p1)
	num, err := libphonenumber.Parse("6502530000", "US")
	t.Log(num)
	p, err := libphonenumber.Parse("+446681800", "") // "+41 44 668 18 00"  "044 668 18 00"
	t.Log(p)
	p3, err := libphonenumber.Parse("+41446681800", "") // "+41 44 668 18 00"  "044 668 18 00"
	t.Log(p3)
	p2, err := libphonenumber.Parse("+371 65 552-336", "LV")
	t.Log(p2)
	p4, err := libphonenumber.Parse("+371 65 552-336", "")
	t.Log(p4)
	p5, err := libphonenumber.Parse("+37165552336", "")
	t.Log(p5)
	p6, err := libphonenumber.Parse("+37165552-336", "")
	t.Log(p6)
	p7, err := libphonenumber.Parse("13312341234", "")
	t.Log(p7)

	t.Log(p1.GetCountryCodeSource())
}

func TestPhoneDomestic(t *testing.T) {
	//p1, _ := libphonenumber.Parse("+8613312341234", "")
	//p7, _ := libphonenumber.Parse("13312341234", "")
	//p4, _ := libphonenumber.Parse("+371 65 552-336", "")
	t.Log(checkPhoneIsDomestic("+8613312341234"))
	t.Log(checkPhoneIsDomestic("+8613312341234"))
	t.Log(checkPhoneIsDomestic("+371 65 552-336"))

}

func checkPhoneIsDomestic(phone string) bool {
	p, _ := libphonenumber.Parse(phone, "")
	if p == nil || p.CountryCode == nil {
		return true
	}
	if *p.CountryCode == 86 {
		return true
	}
	return false
}
