package ebcdic

import (
	"testing"
	"unicode"
)

func TestScanWords(t *testing.T) {

	tests := []struct {
		data    []byte
		atEOF   bool
		advance int
		want    string
	}{
		{
			[]byte("\x40\x40hello\x40world"), false, 8, "hello",
		},
		{
			[]byte("OMG\x40world"), false, 4, "OMG",
		},
		{
			[]byte("\x40OMG"), true, 4, "OMG",
		},
	}

	for _, test := range tests {
		adv, token, _ := ScanWords(test.data, test.atEOF)

		if string(token) != test.want {
			t.Errorf("ScanWords, returns wrong token: got(%s) want(%s)\n", token, test.want)
		}

		if adv != test.advance {
			t.Errorf("ScanWords returns wrong advance: got(%d) want(%d)\n", adv, test.advance)
		}
	}
}

func TestIsSpace(t *testing.T) {
	tests := []struct {
		r    rune
		want bool
	}{
		{' ', true},
		{'\u0085', true},
		{'t', false},
	}

	for _, test := range tests {
		if isSpace(test.r) != unicode.IsSpace(test.r) {
			t.Errorf("isSpace implementation mismathed %v", test.r)
		}
	}

}
