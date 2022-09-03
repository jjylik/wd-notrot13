package affine

import (
	"testing"
)

func TestDecipherAscii(t *testing.T) {
	type test struct {
		input string
		want  string
	}

	tests := []test{
		{input: "OCDNDNVN", want: "THISISAS"},
		{input: "OCDN DNVN\n", want: "THISISAS"},
	}

	for _, tt := range tests {
		got, err := Decrypt(tt.input)
		if got != tt.want || err != nil {
			t.Errorf("Decrypt(%q) = %q, want %q", tt.input, got, tt.want)
		}
	}
}

func TestDecipherNonAscii(t *testing.T) {
	got, err := Decrypt("😎")
	if err == nil {
		t.Errorf("Should return err but got %s", got)
	}
}
