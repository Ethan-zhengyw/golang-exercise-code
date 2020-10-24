package uc

import "testing"

type ucTest struct {
	in, out string
}

var ucTests = []ucTest{
	ucTest{"abc", "ABC"},
	ucTest{"a-bc", "A-BC"},
	// ucTest{"a-bc", "A-Bc"},
}

func TestUC(t *testing.T) {
	for _, item := range ucTests {
		uc := UpperCase(item.in)
		if uc != item.out {
			t.Errorf("Upper case of %v must be %v, got %v", item.in, item.out, uc)
		}
	}
}
