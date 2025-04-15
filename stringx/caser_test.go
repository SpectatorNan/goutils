package stringx

import "testing"

func TestTitle(t *testing.T) {
	str := "COMPLETED"
	t.Logf("Title(%s) = %s", str, Title(str))

	str = "Completed_Funds_Held"
	upper := Upper(str)
	t.Logf("Upper(%s) = %s", str, upper)
	t.Logf("Title(%s) = %s", upper, Title(upper))
	t.Logf("Lower(%s) = %s", upper, Lower(upper))
	t.Logf("Title(%s) = %s", Lower(upper), Title(Lower(upper)))
}
