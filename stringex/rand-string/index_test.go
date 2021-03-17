package randstring

import "testing"

func Test_Create(t *testing.T) {
	size := 10
	str := Create(size)
	if len(str) != size {
		t.Error(str)
	}
}

func Test_Create_100(t *testing.T) {
	size := 100
	str := Create(size)
	if len(str) == size {
		t.Error(str)
	}
}
