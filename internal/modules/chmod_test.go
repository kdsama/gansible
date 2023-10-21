package modules

import "testing"

func TestNewMode(t *testing.T) {

	want := "chmod 0644 /path/to/somewhere"
	got := NewMode("0644", "/path/to/somewhere")
	if want != got {
		t.Errorf("Wanted %v but got %v", want, got)
	}

}
