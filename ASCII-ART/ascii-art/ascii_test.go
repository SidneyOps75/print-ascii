package ascii

import (
	"testing"
)

func TestIsPrintableAscii(t *testing.T) {
	got, err := IsPrintableAscii("h世界")
	want := false

	if got != want || err == nil {
		t.Errorf("got %v, wanted %v with an error", got, want)
	}
}
