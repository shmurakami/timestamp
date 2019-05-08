package timestamp

import "testing"

func TestDate(t *testing.T) {
	s := Source{Time: 1557277200}
	d, _ := s.Date()
	expect := "2019-05-08T10:00:00+09:00"
	if d != expect {
		t.Fatalf("not match\n%s\n%s", d, expect)
	}
}
